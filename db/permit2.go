package db

import (
	"fmt"
	"math/rand/v2"

	"universal-swap/dto"

	"github.com/holiman/uint256"
)

// permit2.go integrates the structs necessary to interact with the Permit2 Solidity contract:
// https://github.com/dragonfly-xyz/useful-solidity-patterns/tree/main/patterns/permit2

type Permit2 struct {
	Permit2Message
	TransferDetails
	Owner     string `json:"owner"`
	Signature []byte `json:"signature"`
}

type Permit2Message struct {
	Permitted TokenPermissions `json:"permitted"`
	Nonce     uint256.Int      `json:"nonce"`    // A unique, consumable number to identify this permit.
	Deadline  uint64           `json:"deadline"` // The latest possible block timestamp for when this permit is valid.
}

type TokenPermissions struct {
	Token  string  `json:"token"`  // Address of the token to be transferred.
	Amount float64 `json:"amount"` // Maximum amount that can be transferred when consuming this permit.
}

type TransferDetails struct {
	To              string  `json:"to"`              // Who receives the permitted token.
	RequestedAmount float64 `json:"requestedAmount"` // How much should be transferred. This can be less than signed in TokenPermissions
}

// The results of BuildPermit2Message will allow interaction with function permitTransferFrom(...)
func BuildPermit2Message(dto dto.CreateOrder, currentBlock uint64, userAddress string, signature []byte) (Permit2, error) {

	message := Permit2Message{
		Permitted: TokenPermissions{
			Token:  dto.SwapToken.Val,
			Amount: dto.InitialPrice.Val,
		},
		Nonce:    *uint256.NewInt(rand.Uint64N(100)),
		Deadline: currentBlock + dto.DecayOffset.Val + dto.DecayDuration.Val,
	}

	details := TransferDetails{
		To:              userAddress,
		RequestedAmount: dto.MinPrice.Val,
	}

	permit := Permit2{
		Permit2Message:  message,
		TransferDetails: details,
		Owner:           userAddress,
		Signature:       signature,
	}

	return permit, nil
}

// Consume a permit2 message and transfer tokens.
func (ex *Exchange) SendPermit(orderId string, permit Permit2) error {

	fmt.Println()
	fmt.Println("Sending signed Permit2...")
	fmt.Println("Owner:", permit.Owner)         // Signed the permit and also holds the tokens (msg.sender).
	fmt.Println("Signature:", permit.Signature) // The EIP-712 signature for the permit2 message, signed by owner.

	// In a production backend we'd send these params to Solidity contract permitTransferFrom()
	// Instead we'll save them in our mock permit repo:
	ex.permits[orderId] = permit

	return nil
}
