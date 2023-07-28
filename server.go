package main

import (
	"api-pressure/controller"
	router "api-pressure/http"
	"fmt"
	"net/http"
)

var (
	pressureController controller.PressureController = controller.NewPressureController()
	httpRouter         router.Router                 = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.POST("/pressures", pressureController.AddPressure)
	httpRouter.GET("/pressures", pressureController.GetPressures)
	httpRouter.SERVE(port)
}
