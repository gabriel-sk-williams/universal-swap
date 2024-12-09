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
	BTC  float64 `json:"btc"`
	SOL  float64 `json:"sol"`
	XRP  float64 `json:"xrp"`
	DOGE float64 `json:"doge"`
	DOT  float64 `json:"dot"`
	NEAR float64 `json:"near"`
	LTC  float64 `json:"ltc"`
	ADA  float64 `json:"ada"`
	BCH  float64 `json:"bch"`
	ALGO float64 `json:"algo"`
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

func (mr *Merchant) ListAssets(ctx context.Context) (Assets, error) {
	return mr.Assets, nil
}

// Sequences of events for issuance of Universal token

// 1. The merchant initiates the issuance process by sending a transaction to the token contract.
// This transaction signals the merchant's intent to send underlying assets equivalent to
// the desired number of Universal tokens. The merchant specifies the blockchain and address
// where they want to receive these tokens

// 2. The merchant transfers the corresponding underlying assets to the custodian.
// The value of these assets should match the value of the Universal tokens they wish to issue

// 3. Upon receiving the underlying assets, the network verifies the received amount. Once the verification is
// successful, the network creates a transaction to mint the new Universal tokens. The specified number of tokens (n)
// are then issued to the merchant's address on their desired blockchain.

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (mr *Merchant) MintTokens(ctx context.Context, dto dto.MintTokens) error {
	mr.RLock()
	defer mr.RUnlock()

	fmt.Println("minting tokens")
	// token, exists := mr.tokens[orderId]

	return nil
}

// To redeem Universal tokens, the merchant must call the ‘burn’ function in the contract,
// specifying the amount of tokens to be burned. When this function is called, the
// specified amount is deducted from the merchant’s Universal token balance on-chain,
// and the overall supply of Universal tokens is reduced accordingly.

// curl -X GET localhost:5000/fill/3b20e878-922b-4458-9d4e-bee3eb5e5848
func (mr *Merchant) BurnTokens(ctx context.Context, dto dto.BurnTokens) error {
	mr.RLock()
	defer mr.RUnlock()

	fmt.Println("burning tokens")
	// token, exists := mr.tokens[orderId]

	return nil
}
