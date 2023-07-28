package controller

import (
	"api-pressure/domain"
	"api-pressure/errors"
	"api-pressure/service"
	"encoding/json"
	"net/http"
)

type PressureController interface {
	AddPressure(resp http.ResponseWriter, req *http.Request)
	GetPressures(resp http.ResponseWriter, req *http.Request)
}

var (
	pressureService service.PressureService = service.NewPressureService()
)

type controller struct{}

func NewPressureController() PressureController {
	return &controller{}
}

func (*controller) AddPressure(resp http.ResponseWriter, req *http.Request) {
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

func (*controller) GetPressures(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	pressures, err := pressureService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error Get Pressure"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pressures)
}
