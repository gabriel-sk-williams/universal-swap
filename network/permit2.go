package network

import (
	"math/rand/v2"

	"github.com/holiman/uint256"
)

// permit2.go integrates the structs necessary to interact with the Permit2 Solidity contract:
// https://github.com/dragonfly-xyz/useful-solidity-patterns/tree/main/patterns/permit2

type Permit2Message struct {
	Permitted TokenPermissions
	Nonce     uint256.Int // A unique, consumable number to identify this permit.
	Deadline  uint64      // The latest possible block timestamp for when this permit is valid.
}

type TokenPermissions struct {
	Token  string  // Address of the token to be transferred.
	Amount float64 // Maximum amount that can be transferred when consuming this permit.
}

type TransferDetails struct {
	To              string  // Who receives the permitted token.
	RequestedAmount float64 // How much should be transferred. This can be less than signed in TokenPermissions
}

func BuildPermit2Message(tokenAddress string, maxAmount float64, address string, amount float64) (Permit2Message, TransferDetails, error) {

	nonce := *uint256.NewInt(rand.Uint64N(100))
	// deadline
	var deadline uint64 = 1000

	permit := Permit2Message{
		Permitted: TokenPermissions{
			Token:  tokenAddress,
			Amount: maxAmount,
		},
		Nonce:    nonce,
		Deadline: deadline,
	}

	details := TransferDetails{
		To: address,
	}

	return permit, details, nil
}

// owner - Who signed the permit and also holds the tokens (msg.sender).

// signature - The corresponding EIP-712 signature for the permit2 message, signed by owner.
// If the recovered address from signature verification does not match owner, the call will fail.

// Consume a permit2 message and transfer tokens.
func PermitTransferFrom(permit Permit2Message, transferDetails TransferDetails, owner string, signature []byte) error {

	return nil
}
