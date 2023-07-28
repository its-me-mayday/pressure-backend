package main

import (
	"api-pressure/controller"
	router "api-pressure/http"
	"api-pressure/service"
	"fmt"
	"net/http"
)

var (
	pressureService    service.PressureService       = service.NewPressureService()
	pressureController controller.PressureController = controller.NewPressureController(pressureService)
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
