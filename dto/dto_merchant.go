package dto

import "goyave.dev/goyave/v5/util/typeutil"

type MintTokens struct {
	Type   typeutil.Undefined[string]  `json:"string"`
	Amount typeutil.Undefined[float64] `json:"float64"`
}

type BurnTokens struct {
	Type   typeutil.Undefined[string]  `json:"string"`
	Amount typeutil.Undefined[float64] `json:"float64"`
}
