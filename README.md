# universal-swap
A minimal UniswapX exchange with uAsset tokens.

This is a backend application simulating an exchange where users can submit Exclusive Dutch Orders, which allow the specification of the maximum and minimum outputs they are willing to receive in a trade over a certain time period. During this period (specified in blocks), the price will slowly decay until the minimum is price is reached.

When calling `CreateOrder` the backend will build an Order from a constructed json provided by the user and submit it to a mockup of the exchange. The server also simulates the provision of a signature by the frontend user to create suitable parameters for interacting with Uniswap's Permit2 contract, which enables users to post and Order and provide a means faciliting transfer without additional permissions.

When calling `FillOrder` the server simulates a Universal Authorized Merchant fulfilling an Order and clearing its entry from the Reactor.

## Getting Started

This server uses Go version 1.22.5

Run `go run main.go` in your project's directory to start the server. The results of server interactions will print to the terminal; or open your browser to `http://localhost:5000`.

The available routes are as follows:
```go
router.Get("/order", ctrl.ListOrders)
router.Post("/order", ctrl.CreateOrder) // creates an Order and submits a Permit2
router.Get("/order/{orderId}", ctrl.GetOrder)
router.Post("/fill/{orderId}", ctrl.FillOrder) // fulfills the given Order
router.Delete("/order/{orderId}", ctrl.DeleteOrder)
router.Get("/permit/{orderId}", ctrl.GetPermitByOrderId)
```

## Interacting with the Server

There is no frontend built for this application. Using a different terminal, users may cURL as provided in the comments of `controller/controls.go` or below in this README:

To start I recommend placing an initial Order to the server:

Windows
```
curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d "{\"Token\": \"BTC\", \"TokenAmount\": 0.0064, \"DecayOffset\": 0, \"DecayDuration\": 300, \"SwapToken\": \"USDC\", \"InitialPrice\": 636.00, \"MinPrice\": 600.00}"
```

Linux
```
curl -X POST http://localhost:5000/order -H "Content-Type: application/json" -d "{"Token": "BTC", "TokenAmount": 0.0064, "DecayOffset": 0, "DecayDuration": 300, "SwapToken": "USDC", "InitialPrice": 636.00, "MinPrice": 600.00}"
```

### Directory structure

```
.
├── controller
│   └── controls.go          // asdf
│   └── error.go             // asdf
├── db
│   └── exchange.go          // asdf
│   └── order.go             // asdf
│   └── permit2.go           // asdf
├── dto
│   └── dto_order.go         // asdf
├── network
│   └── crypto.go            // asdf
├── route
│   └── route.go             // Static resources
│
├── .gitignore
├── config.json              // config for local development
├── go.mod
├── go.sum
└── main.go                  // Application entrypoint
```

```

```

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then open your browser to `http://localhost:8080`.
