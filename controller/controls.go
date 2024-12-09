package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"universal-swap/db"
	"universal-swap/dto"
	"universal-swap/network"

	"github.com/ethereum/go-ethereum/ethclient"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Controller struct {
	goyave.Component
	EX *db.Exchange
	CL *ethclient.Client
}

func NewController(server *goyave.Server, exchange *db.Exchange, client *ethclient.Client) *Controller {
	ctrl := &Controller{
		EX: exchange,
		CL: client,
	}
	ctrl.Init(server)
	return ctrl
}

// curl -X GET localhost:5000/order
func (c *Controller) ListOrders(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	results, err := c.EX.ListOrders(ctx)

	if err == nil {
		res.JSON(http.StatusOK, results)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d "{\"Token\": \"BTC\", \"TokenAmount\": 0.0064, \"DecayOffset\": 0, \"DecayDuration\": 300, \"SwapToken\": \"USDC\", \"InitialPrice\": 636.00, \"MinPrice\": 600.00}"
func (c *Controller) CreateOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	// Simulate the components needed to build the Order and sign
	privateKey, userAddress := network.GenerateUser()

	// Receive and convert json from frontend into out Data Transfer Object
	dto, err := typeutil.Convert[dto.CreateOrder](req.Data)
	check(err)

	// Get current block time
	currentBlock, err := c.CL.BlockNumber(ctx)
	check(err)

	// Create the order
	order, err := c.EX.CreateOrder(ctx, dto, currentBlock, userAddress)
	check(err)

	fmt.Println()
	fmt.Printf("New order posted from wallet address: %s \n", userAddress)
	fmt.Printf("Token: %f %s", order.Input.Amount, order.Input.Token)
	fmt.Printf("Swap: %f->%f %s", order.Output.StartAmount, order.Output.EndAmount, order.Output.Token)

	// Convert the order to byte array
	orderBytes, err := json.Marshal(order)
	if err != nil {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}

	// Use the randomly generated private key to sign the order
	signature, err := network.SignOrder(privateKey, orderBytes)
	check(err)

	// Build our params for interaction with the UniswapX Permit2 contract
	permit, err := db.BuildPermit2Message(dto, currentBlock, userAddress, signature)
	check(err)

	// Simulate sending the permit
	err = c.EX.SendPermit(order.Info.ID, permit)
	check(err)

	if err == nil {
		res.JSON(http.StatusCreated, order)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl localhost:5000/order/aa44c41d-0bd2-4baf-a587-11f984f6910b
func (c *Controller) GetOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	orderId := req.RouteParams["orderId"]
	result, exists := c.EX.GetOrder(ctx, orderId)

	if exists {
		res.JSON(http.StatusOK, result)
	} else {
		res.Status(http.StatusNotFound)
		res.Error(exists)
	}
}

// curl -X POST localhost:5000/fill/0d191cb2-5594-4dcc-aa5e-bd42334739f6
func (c *Controller) FillOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	// Get current block time
	currentBlock, err := c.CL.BlockNumber(ctx)
	check(err)

	orderId := req.RouteParams["orderId"]
	err = c.EX.FillOrder(ctx, orderId, currentBlock)

	if err == nil {
		res.JSON(http.StatusOK, fmt.Sprintf("Filled order: %s", orderId))
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl -X DELETE localhost:5000/order/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (c *Controller) DeleteOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	orderId := req.RouteParams["orderId"]
	err := c.EX.DeleteOrder(ctx, orderId)
	if err != nil {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}

	err = c.EX.DeletePermit(ctx, orderId)

	if err == nil {
		res.JSON(http.StatusOK, fmt.Sprintf("Deleted order: %s", orderId))
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl localhost:5000/permit/0d191cb2-5594-4dcc-aa5e-bd42334739f6
func (c *Controller) GetPermitByOrderId(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	orderId := req.RouteParams["orderId"]
	result, exists := c.EX.GetPermit(ctx, orderId)

	if exists {
		res.JSON(http.StatusOK, result)
	} else {
		res.Status(http.StatusNotFound)
		res.Error(exists)
	}
}
