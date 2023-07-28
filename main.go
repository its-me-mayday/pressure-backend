package main

import (
	router "api-pressure/http"
	"fmt"
	"net/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.SERVE(port)
}
