package dto

import "goyave.dev/goyave/v5/util/typeutil"

type CreateOrder struct {
	Token         typeutil.Undefined[string]  `json:"token"`
	TokenAmount   typeutil.Undefined[float64] `json:"tokenAmount"`
	DecayOffset   typeutil.Undefined[uint64]  `json:"decayOffset"`
	DecayDuration typeutil.Undefined[uint64]  `json:"decayDuration"`
	SwapToken     typeutil.Undefined[string]  `json:"swapToken"`
	InitialPrice  typeutil.Undefined[float64] `json:"initialPrice"`
	MinPrice      typeutil.Undefined[float64] `json:"minPrice"`
}
