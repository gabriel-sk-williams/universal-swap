package db

import "fmt"

// ExclusiveDutchOrder via UniswapX:
// https://github.com/Uniswap/UniswapX/blob/main/src/lib/ExclusiveDutchOrderLib.sol
type ExclusiveDutchOrder struct {
	Info           OrderInfo   `json:"info"`           // Generic order information
	DecayStartTime uint64      `json:"decayStartTime"` // The time at which the DutchOutputs start decaying
	DecayEndTime   uint64      `json:"decayEndTime"`   // The time at which price becomes static
	Input          DutchInput  `json:"input"`          // The tokens that the swapper will provide when settling the order
	Output         DutchOutput `json:"outputs"`        // The tokens that must be received to satisfy the order
	// ExclusiveFiller        string                   // omitted for this mockup
	// ExclusivityOverrideBps uint64                   // omitted -> btw 100 basis points = 1%
}

// OrderInfo via UniswapX:
// https://github.com/Uniswap/UniswapX/blob/main/src/base/ReactorStructs.sol
type OrderInfo struct {
	ID       string `json:"id"`       // OrderId for tracking in mock db
	Reactor  string `json:"reactor"`  // The address of the reactor that this order is targeting
	Swapper  string `json:"swapper"`  // The address of the user which created the order
	Deadline uint64 `json:"deadline"` // The timestamp after which this order is no longer valid
	// Nonce uint64                        // omitted -> Allows for signature replay protection and cancellation
	// AdditionalValidationContract string // omitted for this mockup
	// AdditionalValidationData     string // omitted for this mockup
}

// DutchInput via UniswapX:
// https://github.com/Uniswap/UniswapX/blob/main/src/lib/DutchOrderLib.sol
type DutchInput struct {
	Token  string  `json:"token"`  // The ERC20 token address
	Amount float64 `json:"amount"` // The amount of tokens
}

// DutchOutput via UniswapX:
// https://github.com/Uniswap/UniswapX/blob/main/src/lib/DutchOrderLib.sol
type DutchOutput struct {
	Token       string  `json:"token"`       // The ERC20 token address (or native ETH address)
	StartAmount float64 `json:"startAmount"` // The amount of tokens at the start of the time period
	EndAmount   float64 `json:"endAmount"`   // The amount of tokens at the end of the time period
	Recipient   string  `json:"recipient"`   // The address who must receive the tokens to satisfy the order
}

// LinearPriceDecay calculates price using linear decay
// currentPrice := order.LinearPriceDecay(time.Now().Add(30 * time.Minute))
func (do *ExclusiveDutchOrder) LinearPriceDecay(currentBlock uint64) float64 {

	fmt.Println("Decay Start:", do.DecayStartTime)
	fmt.Println("Decay End:", do.DecayEndTime)
	fmt.Println("Decay Deadline:", do.Info.Deadline)

	// if decay has not started yet, return startAmount
	if do.DecayStartTime > currentBlock {
		return do.Output.StartAmount
	}

	// Calculate time elapsed
	timeElapsed := currentBlock - do.DecayStartTime
	fmt.Println("Blocks Elapsed:", timeElapsed)

	decayDuration := do.DecayEndTime - do.DecayStartTime

	// Prevent exceeding decay duration
	if timeElapsed > decayDuration {
		return do.Output.EndAmount
	}

	// Calculate total price drop
	totalPriceDrop := do.Output.StartAmount - do.Output.EndAmount

	// Calculate current price drop based on elapsed time
	currentPriceDrop := (float64(timeElapsed) / float64(decayDuration)) * totalPriceDrop

	// Calculate current price
	currentPrice := do.Output.StartAmount - currentPriceDrop

	// Ensure price doesn't go below minimum
	if currentPrice < do.Output.EndAmount {
		return do.Output.EndAmount
	}

	return currentPrice
}
