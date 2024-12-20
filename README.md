# universal-swap
A minimal UniswapX exchange with uAsset tokens.

This is a backend application simulating an exchange where users can swap coins by submitting Exclusive Dutch Orders, which allow the specification of the maximum and minimum outputs they are willing to receive in a trade over a certain time period. During this period (specified in blocks), the price will slowly decay until the minimum price is reached.

When calling `CreateOrder` the backend will build an Order from a constructed json and submit it to a mockup of a Uniswap Dutch Reactor. The server also simulates the provision of a signature for use with Uniswap's Permit2 contract, faciliting token transfer without additional permissions.

When calling `FillOrder` the server simulates a Universal Authorized Merchant fulfilling an Order and clearing its entry from the Reactor.

## Getting Started

This project uses Go version 1.22.5

Enter `go run .` or `go run main.go` in the project's directory to start the server. The results of server interactions will print to the terminal. You can also open your browser to `http://localhost:5000`.

The available routes are as follows:
```go
router.Get("/order", ctrl.ListOrders)
router.Post("/order", ctrl.CreateOrder) // create an Order and submit a Permit2
router.Get("/order/{orderId}", ctrl.GetOrder)
router.Post("/fill/{orderId}", ctrl.FillOrder) // fulfill a given Order
router.Delete("/order/{orderId}", ctrl.DeleteOrder)
router.Get("/permit/{orderId}", ctrl.GetPermitByOrderId)
```

## Interacting with the Server

There is no frontend for this application.

To start, I recommend sending an Order to the server via cURL:

Windows:
```
curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d "{\"Token\": \"BTC\", \"TokenAmount\": 0.0064, \"DecayOffset\": 0, \"DecayDuration\": 300, \"SwapToken\": \"USDC\", \"InitialPrice\": 636.00, \"MinPrice\": 600.00}"
```

Linux:
```
curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d '{"Token": "BTC", "TokenAmount": 0.0064, "DecayOffset": 0, "DecayDuration": 300, "SwapToken": "USDC", "InitialPrice": 636.00, "MinPrice": 600.00}'
```

Placing the Order will store it in memory and make it accessible with a GET to `/order` or `order/{orderId}`. To fill the Order, copy the id and use it in your POST to `/fill/{orderId}`, e.g.:

`curl -X POST http://localhost:5000/order/b3b0bd43-d7ba-47a8-8b4d-91fc058cc9e2`

The server will calculate price decay based on the parameters specified, and the number of blocks elapsed on the Base L2 Network. You should receive a confirmation on the server terminal:

```
Order b3b0bd43-d7ba-47a8-8b4d-91fc058cc9e2 fulfilled:
0.006400 BTC swapped at price: 613.680000 USDC
```

### Directory structure

```
.
├── controller
│   └── controls.go          // Controller implementation
│   └── error.go             // Custom functions for error handling
├── db
│   └── exchange.go          // Exchange definitions and methods
│   └── order.go             // Exclusive Dutch Order definitions and methods
│   └── permit2.go           // Permit2 definitions and methods
├── dto
│   └── dto_order.go         // Defines the CreateOrder data transfer object
├── network
│   └── crypto.go            // Cryptographic utility functions
├── route
│   └── route.go             // Route and controller registration
│
├── .gitignore
├── config.json              // Config file for local development
├── go.mod
├── go.sum
└── main.go                  // Application entrypoint
```


### Concluding Thoughts

I would have liked to give more attention to uAssets and the Universal Merchant. In early iterations I included a "Merchant" struct in the controller that would hold mock crypto assets, use them to mint and burn Universal Tokens, and provide liquidity to the Exchange. However, since most of these actions would be occurring on-chain (and therefore off-server), the best I could have done is mock up a controlled entity that 1) would not technically be using my routes and 2) is not actually subject to on-chain limitations.

Instead I decided to focus on preparing frontend requests for eventual use on-chain. The bulk of my efforts went therefore into configuring and signing Dutch Orders that would interact with the UniswapX DutchReactor and Permit2 Solidity contracts:

https://github.com/Uniswap/UniswapX/blob/main/src/lib/ExclusiveDutchOrderLib.sol
https://github.com/Uniswap/UniswapX/blob/main/src/base/ReactorStructs.sol
https://github.com/Uniswap/UniswapX/blob/main/src/lib/DutchOrderLib.sol

https://github.com/dragonfly-xyz/useful-solidity-patterns/tree/main/patterns/permit2


