package db

import (
	"context"
	"sync"
	"time"
	"universal-swap/dto"
)

type Exchange struct {
	sync.RWMutex
	orders map[string]ExclusiveDutchOrder
}

func NewExchange() *Exchange {
	return &Exchange{
		orders: make(map[string]ExclusiveDutchOrder),
	}
}

func (ex *Exchange) ListOrders(ctx context.Context) (map[string]ExclusiveDutchOrder, error) {
	return ex.orders, nil
}

func (ex *Exchange) CreateOrder(ctx context.Context, id string, dto dto.CreateOrder) (ExclusiveDutchOrder, error) {
	ex.Lock()
	defer ex.Unlock()

	order := ExclusiveDutchOrder{
		InitialPrice: dto.InitialPrice.Val,
		MinPrice:     dto.MinPrice.Val,
		MaxPrice:     dto.MaxPrice.Val,
		StartTime:    time.Now(),
		// DecayDuration: time.Hour,
	}

	ex.orders[id] = order

	return order, nil
}

func (ex *Exchange) GetOrder(ctx context.Context, orderId string) (ExclusiveDutchOrder, bool) {
	ex.RLock()
	defer ex.RUnlock()

	order, exists := ex.orders[orderId]
	return order, exists
}

func (ex *Exchange) DeleteOrder(ctx context.Context, orderId string) error {
	ex.RLock()
	defer ex.RUnlock()

	// delete() is simply a no-op if the entry does not exist
	delete(ex.orders, orderId)
	return nil
}
