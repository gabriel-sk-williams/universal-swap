package route

import (
	"fmt"
	"log"
	"net/http"
	"universal-swap/controller"
	"universal-swap/db"

	"github.com/ethereum/go-ethereum/ethclient"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

func Register(server *goyave.Server, router *goyave.Router) {

	env := "local"
	fmt.Printf("Environment: %s", env)

	orderBook := db.NewExchange()
	client := dialClient()

	{
		corsOptions := cors.Default()
		origins := []string{"*"}
		corsOptions.AllowCredentials = true
		corsOptions.AllowedOrigins = origins

		router.CORS(corsOptions)
		router.GlobalMiddleware(&parse.Middleware{})

		ctrl := controller.NewController(server, orderBook, client)

		// UNPROTECTED ROUTES
		router.Get("/", Greeting)

		// EXCHANGE ROUTES
		router.Get("/order", ctrl.ListOrders)
		router.Post("/order", ctrl.CreateOrder)
		router.Get("/order/{orderId}", ctrl.GetOrder)
		router.Post("/fill/{orderId}", ctrl.FillOrder)
		router.Delete("/order/{orderId}", ctrl.DeleteOrder)
		router.Get("/permit/{orderId}", ctrl.GetPermitByOrderId)
	}
}

func dialClient() *ethclient.Client {
	client, err := ethclient.Dial("https://base-mainnet.g.alchemy.com/v2/FYWP3aqwxz3E-jAT8DKWHKMRdYBdPqx-")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum node: %v", err)
	}
	return client
}

// curl http://localhost:5000
func Greeting(res *goyave.Response, req *goyave.Request) {
	res.String(http.StatusOK, "Oh? A fellow coin swapper!?")
}
