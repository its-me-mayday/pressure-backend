package main

import (
	"api-pressure/controller"
	router "api-pressure/http"
	"api-pressure/repository"
	"api-pressure/service"
	"fmt"
	"net/http"
	"os"
)

var (
	Repository         repository.Repository         = repository.NewMockRepository()
	pressureService    service.PressureService       = service.NewPressureService(Repository)
	pressureController controller.PressureController = controller.NewPressureController(pressureService)
	httpRouter         router.Router                 = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.POST("/pressures", pressureController.AddPressure)
	httpRouter.GET("/pressures", pressureController.GetPressures)
	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
