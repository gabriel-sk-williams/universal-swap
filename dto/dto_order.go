package dto

import "goyave.dev/goyave/v5/util/typeutil"

// add types for BTC, Ethereum, Solana
type CreateOrder struct {
	InitialPrice typeutil.Undefined[float64] `json:"initialPrice"`
	MinPrice     typeutil.Undefined[float64] `json:"minPrice"`
	MaxPrice     typeutil.Undefined[float64] `json:"maxPrice"`
}
