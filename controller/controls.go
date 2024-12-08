package controller

import (
	"context"
	"fmt"
	"net/http"
	"universal-swap/db"
	"universal-swap/dto"
	"universal-swap/network"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Controller struct {
	goyave.Component
	EX *db.Exchange
	MR *db.Merchant
	CL *ethclient.Client
}

func NewController(server *goyave.Server, exchange *db.Exchange, merchant *db.Merchant, client *ethclient.Client) *Controller {
	ctrl := &Controller{
		EX: exchange,
		MR: merchant,
		CL: client,
	}
	ctrl.Init(server)
	return ctrl
}

func (c *Controller) GetStatus(res *goyave.Response, req *goyave.Request) {
	res.String(http.StatusOK, "Status OK")
}

// curl -X GET localhost:5000/block
func (c *Controller) GetBlockTime(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()
	result, err := c.CL.BlockNumber(ctx)

	if err == nil {
		res.JSON(http.StatusOK, result)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
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

/*
curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d "{\"InitialPrice\": 100.00, \"MaxPrice\": 120.00, \"MinPrice\": 80.00}"
*/

func (c *Controller) CreateOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	dto, err := typeutil.Convert[dto.CreateOrder](req.Data)
	check(err)

	id := uuid.New().String()[:8] // shortened uuid for readability
	order, err := c.EX.CreateOrder(ctx, id, dto)

	// generates a random user to sign the order and derive a wallet address
	signature, publicKey, err := network.SignOrder("Hello!")
	ethereumAddress := network.PublicKeyToEthereumAddress(publicKey)

	fmt.Println()
	fmt.Printf("New order posted from wallet address %s \n", ethereumAddress)
	fmt.Printf("Signature: %s \n", signature)

	if err == nil {
		res.JSON(http.StatusCreated, order)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl localhost:5000/order/3b20e878-922b-4458-9d4e-bee3eb5e5848
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

// curl -X POST localhost:5000/order/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (c *Controller) DeleteOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	orderId := req.RouteParams["orderId"]
	err := c.EX.DeleteOrder(ctx, orderId)

	if err == nil {
		res.JSON(http.StatusOK, fmt.Sprintf("Deleted order: %s", orderId))
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (c *Controller) FillOrder(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	orderId := req.RouteParams["orderId"]
	err := c.MR.FillOrder(ctx, orderId)

	if err == nil {
		res.JSON(http.StatusOK, fmt.Sprintf("Filled order: %s", orderId))
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

func (c *Controller) MintTokens(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	dto, err := typeutil.Convert[dto.MintTokens](req.Data)
	check(err)

	err = c.MR.MintTokens(ctx, dto)

	if err == nil {
		res.JSON(http.StatusOK, err)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}

func (c *Controller) BurnTokens(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	dto, err := typeutil.Convert[dto.BurnTokens](req.Data)
	check(err)

	err = c.MR.BurnTokens(ctx, dto)

	if err == nil {
		res.JSON(http.StatusOK, err)
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}
