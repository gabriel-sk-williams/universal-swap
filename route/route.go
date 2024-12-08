package route

import (
	"fmt"
	"net/http"
	"universal-swap/controller"
	"universal-swap/db"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

func Register(server *goyave.Server, router *goyave.Router) {

	env := "local"
	fmt.Printf("Environment: %s", env)

	orderBook := db.NewOrderBook()

	{
		corsOptions := cors.Default()
		origins := []string{"*"}
		corsOptions.AllowCredentials = true
		corsOptions.AllowedOrigins = origins

		router.CORS(corsOptions)
		router.GlobalMiddleware(&parse.Middleware{})

		ctrl := controller.NewController(server, orderBook) // driver

		// UNPROTECTED STATUS ROUTES
		router.Get("/", Greeting)
		router.Get("/status", ctrl.GetStatus)

		router.Get("/order", ctrl.ListOrders)
		router.Post("/order", ctrl.CreateOrder)
		router.Get("/order/{orderId}", ctrl.GetOrder)
		router.Delete("/order/{orderId}", ctrl.DeleteOrder)
	}
}

// curl http://localhost:5000
func Greeting(res *goyave.Response, req *goyave.Request) {
	res.String(http.StatusOK, "Oh? A fellow swapper!?")
}
