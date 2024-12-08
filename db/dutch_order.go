package db

import (
	"math"
	"time"
)

// https://github.com/Uniswap/UniswapX/blob/main/src/lib/ExclusiveDutchOrderLib.sol

/*
struct ExclusiveDutchOrder {
    // generic order information
    OrderInfo info;
    // The time at which the DutchOutputs start decaying
    uint256 decayStartTime;
    // The time at which price becomes static
    uint256 decayEndTime;
    // The address who has exclusive rights to the order until decayStartTime
    address exclusiveFiller;
    // The amount in bps that a non-exclusive filler needs to improve the outputs by to be able to fill the order
    uint256 exclusivityOverrideBps;
    // The tokens that the swapper will provide when settling the order
    DutchInput input;
    // The tokens that must be received to satisfy the order
    DutchOutput[] outputs;
}
*/

type ExclusiveDutchOrder struct {
	InitialPrice float64
	MinPrice     float64
	MaxPrice     float64
	StartTime    time.Time

	DecayStartTime time.Duration
	DecayEndTime   time.Duration
}

// LinearPriceDecay calculates price using linear decay
// currentPrice := order.LinearPriceDecay(time.Now().Add(30 * time.Minute))
func (do *ExclusiveDutchOrder) LinearPriceDecay(currentTime time.Time) float64 {
	// Calculate time elapsed
	timeElapsed := currentTime.Sub(do.StartTime)

	decayDuration := do.DecayEndTime - do.DecayStartTime

	// Prevent negative time
	if timeElapsed < 0 {
		timeElapsed = 0
	}

	// Prevent exceeding decay duration
	if timeElapsed > decayDuration {
		return do.MinPrice
	}

	// Calculate total price drop
	totalPriceDrop := do.InitialPrice - do.MinPrice

	// Calculate current price drop based on elapsed time
	currentPriceDrop := (float64(timeElapsed) / float64(decayDuration)) * totalPriceDrop

	// Calculate current price
	currentPrice := do.InitialPrice - currentPriceDrop

	// Ensure price doesn't go below minimum
	if currentPrice < do.MinPrice {
		return do.MinPrice
	}

	return currentPrice
}

// ExponentialPriceDecay calculates price using exponential decay
// exponentialPrice := order.ExponentialPriceDecay(time.Now().Add(30*time.Minute), 0.001)
func (do *ExclusiveDutchOrder) ExponentialPriceDecay(currentTime time.Time, decayRate float64) float64 {
	timeElapsed := currentTime.Sub(do.StartTime).Seconds()

	// Exponential decay formula
	currentPrice := do.InitialPrice * math.Exp(-decayRate*timeElapsed)

	// Ensure price doesn't go below minimum
	if currentPrice < do.MinPrice {
		return do.MinPrice
	}

	return currentPrice
}
