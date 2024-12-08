# universal-swap
A minimal UniswapX exchange with uAsset tokens

A minimal chat application to showcase [Goyave](https://github.com/go-goyave/goyave)'s websocket feature. This project is based on [Gorilla's chat example](https://github.com/gorilla/websocket/tree/master/examples/chat).

**Disclaimer:** This example project cannot be used in a real-life scenario, as you would need to be able to serve clients across multiple instances of the application. This is a typical scenario in cloud environments. The hub in this example could use a PUB/SUB mechanism (for example with [redis](https://redis.io/docs/interact/pubsub/)) to solve this issue.

## Getting Started

### Directory structure

```
.
├── controller
│   └── controls.go          // asdf
│   └── error.go             // asdf
├── db
│   └── book.go              // asdf
│   └── order.go             // asdf
├── dto                      // asdf
│   └── dto_order.go         // asdf
├── route
│   └── route.go             // Static resources
│
├── .gitignore
├── config.json             // config for local development
├── go.mod
├── go.sum
└── main.go                  // Application entrypoint
```

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then open your browser to `http://localhost:8080`.
