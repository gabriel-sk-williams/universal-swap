package db

import (
	"context"
	"fmt"
	"sync"
	"universal-swap/dto"
	"universal-swap/network"

	"github.com/google/uuid"
)

type Exchange struct {
	sync.RWMutex
	address string
	permits map[string]Permit2
	orders  map[string]ExclusiveDutchOrder
}

func NewExchange() *Exchange {
	return &Exchange{
		address: network.GenerateEthereumAddress(),
		permits: make(map[string]Permit2),
		orders:  make(map[string]ExclusiveDutchOrder),
	}
}

func (ex *Exchange) ListOrders(ctx context.Context) (map[string]ExclusiveDutchOrder, error) {
	return ex.orders, nil
}

func (ex *Exchange) CreateOrder(ctx context.Context, dto dto.CreateOrder, currentBlock uint64, userAddress string) (ExclusiveDutchOrder, error) {
	ex.Lock()
	defer ex.Unlock()

	orderId := uuid.New().String()

	order := ExclusiveDutchOrder{
		Info: OrderInfo{
			ID:       orderId,
			Reactor:  ex.address,
			Swapper:  userAddress,
			Deadline: currentBlock + 7200, // roughly 1 hour on Base L2
		},
		DecayStartTime: currentBlock + dto.DecayOffset.Val,
		DecayEndTime:   currentBlock + dto.DecayOffset.Val + dto.DecayDuration.Val,
		Input: DutchInput{
			Token:  dto.Token.Val,
			Amount: dto.TokenAmount.Val,
		},
		Output: DutchOutput{
			Token:       dto.SwapToken.Val,
			StartAmount: dto.InitialPrice.Val,
			EndAmount:   dto.MinPrice.Val,
			Recipient:   userAddress,
		},
	}

	ex.orders[orderId] = order

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

func (ex *Exchange) DeletePermit(ctx context.Context, orderId string) error {
	ex.RLock()
	defer ex.RUnlock()

	// delete() is simply a no-op if the entry does not exist
	delete(ex.permits, orderId)
	return nil
}

func (ex *Exchange) GetPermit(ctx context.Context, orderId string) (Permit2, bool) {
	ex.RLock()
	defer ex.RUnlock()

	permit, exists := ex.permits[orderId]
	return permit, exists
}

func (ex *Exchange) FillOrder(ctx context.Context, orderId string, currentBlock uint64) error {
	ex.RLock()
	defer ex.RUnlock()

	// Simulate a merchant filling the order
	order, exists := ex.orders[orderId]
	if !exists {
		return fmt.Errorf("%q order not found", orderId)
	}

	currentPrice := order.LinearPriceDecay(currentBlock)

	fmt.Println()
	fmt.Printf("Order %s fulfilled: \n", orderId)
	fmt.Printf("%f %s swapped at price: %f %s \n", order.Input.Amount, order.Input.Token, currentPrice, order.Output.Token)

	delete(ex.orders, orderId)
	delete(ex.permits, orderId)

	return nil
}
