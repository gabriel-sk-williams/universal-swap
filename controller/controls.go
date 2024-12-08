package controller

import (
	"context"
	"fmt"
	"net/http"
	"universal-swap/db"
	"universal-swap/dto"

	"github.com/google/uuid"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Controller struct {
	goyave.Component
	OB *db.OrderBook
}

func NewController(server *goyave.Server, orderBook *db.OrderBook) *Controller {
	ctrl := &Controller{
		OB: orderBook,
	}
	ctrl.Init(server)
	return ctrl
}

func (c *Controller) GetStatus(res *goyave.Response, req *goyave.Request) {
	res.String(http.StatusOK, "Status OK")
}

// curl -X GET localhost:5000/order
func (c *Controller) ListOrders(res *goyave.Response, req *goyave.Request) {
	ctx := context.Background()

	results, err := c.OB.ListOrders(ctx)

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

	id := uuid.New().String()
	order, err := c.OB.CreateOrder(ctx, id, dto)

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
	result, exists := c.OB.GetOrder(ctx, orderId)

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
	err := c.OB.DeleteOrder(ctx, orderId)

	if err == nil {
		res.JSON(http.StatusOK, fmt.Sprintf("Deleted order: %s", orderId))
	} else {
		res.Status(http.StatusInternalServerError)
		res.Error(err)
	}
}
