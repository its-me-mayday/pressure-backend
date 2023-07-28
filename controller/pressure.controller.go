package controller

import (
	"api-pressure/domain"
	"api-pressure/errors"
	"api-pressure/service"
	"encoding/json"
	"net/http"
)

type PressureController interface {
	getPressures(resp http.ResponseWriter, req *http.Request)
	addPressure(resp http.ResponseWriter, req *http.Request)
}

var (
	pressureService service.PressureService = service.NewPressureService()
)

func getPressures(resp http.ResponseWriter, req *http.Request) {
	//resp.Header().Set("Content-type", "application/json")
	//result, err := json.Marshal(pressures)
	//if err != nil {
	//	resp.WriteHeader(http.StatusInternalServerError)
	//	resp.Write([]byte(`{"error","Error marshalling"}`))
	//	return
	//}
	//resp.WriteHeader(http.StatusOK)
	//resp.Write(result)
}

func addPressure(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var pressure domain.Pressure
	err := json.NewDecoder(req.Body).Decode(&pressure)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error Add Pressure"})
		return
	}

	pressureService.Create(&pressure)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(pressure)
	resp.Write(result)
}
