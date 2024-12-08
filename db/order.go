package db

import (
	"math"
	"time"
)

type DutchOrder struct {
	InitialPrice  float64
	MinPrice      float64
	MaxPrice      float64
	StartTime     time.Time
	DecayDuration time.Duration
}

// LinearPriceDecay calculates price using linear decay
func (do *DutchOrder) LinearPriceDecay(currentTime time.Time) float64 {
	// Calculate time elapsed
	timeElapsed := currentTime.Sub(do.StartTime)

	// Prevent negative time
	if timeElapsed < 0 {
		timeElapsed = 0
	}

	// Prevent exceeding decay duration
	if timeElapsed > do.DecayDuration {
		return do.MinPrice
	}

	// Calculate total price drop
	totalPriceDrop := do.InitialPrice - do.MinPrice

	// Calculate current price drop based on elapsed time
	currentPriceDrop := (float64(timeElapsed) / float64(do.DecayDuration)) * totalPriceDrop

	// Calculate current price
	currentPrice := do.InitialPrice - currentPriceDrop

	// Ensure price doesn't go below minimum
	if currentPrice < do.MinPrice {
		return do.MinPrice
	}

	return currentPrice
}

// ExponentialPriceDecay calculates price using exponential decay
func (do *DutchOrder) ExponentialPriceDecay(currentTime time.Time, decayRate float64) float64 {
	timeElapsed := currentTime.Sub(do.StartTime).Seconds()

	// Exponential decay formula
	currentPrice := do.InitialPrice * math.Exp(-decayRate*timeElapsed)

	// Ensure price doesn't go below minimum
	if currentPrice < do.MinPrice {
		return do.MinPrice
	}

	return currentPrice
}
