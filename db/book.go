package db

import (
	"context"
	"sync"
	"time"
	"universal-swap/dto"
)

type OrderBook struct {
	sync.RWMutex
	orders map[string]DutchOrder
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		orders: make(map[string]DutchOrder),
	}
}

func (ob *OrderBook) ListOrders(ctx context.Context) (map[string]DutchOrder, error) {
	return ob.orders, nil
}

func (ob *OrderBook) CreateOrder(ctx context.Context, id string, dto dto.CreateOrder) (DutchOrder, error) {
	ob.Lock()
	defer ob.Unlock()

	order := DutchOrder{
		InitialPrice:  dto.InitialPrice.Val,
		MinPrice:      dto.MinPrice.Val,
		MaxPrice:      dto.MaxPrice.Val,
		StartTime:     time.Now(),
		DecayDuration: time.Hour,
	}

	ob.orders[id] = order

	return order, nil
}

func (ob *OrderBook) GetOrder(ctx context.Context, orderId string) (DutchOrder, bool) {
	ob.RLock()
	defer ob.RUnlock()

	order, exists := ob.orders[orderId]
	return order, exists
}

func (ob *OrderBook) DeleteOrder(ctx context.Context, orderId string) error {
	ob.RLock()
	defer ob.RUnlock()

	// delete() is simply a no-op if the entry does not exist
	delete(ob.orders, orderId)
	return nil
}

/*
// Get current price after some time
currentPrice := order.LinearPriceDecay(time.Now().Add(30 * time.Minute))

// Exponential decay example with a specific decay rate
exponentialPrice := order.ExponentialPriceDecay(time.Now().Add(30*time.Minute), 0.001)
*/
