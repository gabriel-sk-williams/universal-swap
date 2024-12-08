package db

import (
	"context"
	"fmt"
	"sync"

	"universal-swap/dto"
)

type Merchant struct {
	sync.RWMutex
	BaseWallet
	Assets
}

type BaseWallet struct {
	Address string
	Tokens  map[string]UniversalToken
}

type UniversalToken struct {
	Name   string
	Amount float64
}

type Assets struct {
	BTC  float64
	SOL  float64
	XRP  float64
	DOGE float64
	DOT  float64
	NEAR float64
	LTC  float64
	ADA  float64
	BCH  float64
	ALGO float64
}

func NewMerchant() *Merchant {
	return &Merchant{
		Assets: Assets{
			BTC:  100.0,
			SOL:  100.0,
			XRP:  100.0,
			DOGE: 100.0,
			DOT:  100.0,
			NEAR: 100.0,
			LTC:  100.0,
			ADA:  100.0,
			BCH:  100.0,
			ALGO: 100.0,
		},
	}
}

/*
Sequences of events for issuance of Universal token
1. The merchant initiates the issuance process by sending a transaction to the token contract. This transaction
signals the merchant's intent to send underlying assets equivalent to the desired number of Universal tokens. The
merchant specifies the blockchain and address where they want to receive these tokensÊ
2. The merchant transfers the corresponding underlying assets to the custodian. The value of these assets should
match the value of the Universal tokens they wish to issueÊ
3. Upon receiving the underlying assets, the network verifies the received amount. Once the verification is
successful, the network creates a transaction to mint the new Universal tokens. The specified number of tokens (n)
are then issued to the merchant's address on their desired blockchain.
*/

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (mr *Merchant) FillOrder(ctx context.Context, orderId string) error {
	mr.RLock()
	defer mr.RUnlock()

	// token, exists := mr.tokens[orderId]

	return nil
}

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (mr *Merchant) MintTokens(ctx context.Context, dto dto.MintTokens) error {
	mr.RLock()
	defer mr.RUnlock()

	fmt.Println("minting tokens")
	// token, exists := mr.tokens[orderId]

	return nil
}

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (mr *Merchant) BurnTokens(ctx context.Context, dto dto.BurnTokens) error {
	mr.RLock()
	defer mr.RUnlock()

	fmt.Println("burning tokens")
	// token, exists := mr.tokens[orderId]

	return nil
}
