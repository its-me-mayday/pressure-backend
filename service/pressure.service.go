package service

import (
	"api-pressure/domain"
	"api-pressure/repository"

	"github.com/google/uuid"
)

type PressureService interface {
	//Validate(pressure *domain.Pressure) bool
	Create(pressure *domain.Pressure) (*domain.Pressure, error)
	//FindAll() ([]domain.Pressure, error)
}

type service struct{}

var (
	repo repository.PressureRepository = repository.NewPressureRepository()
)

func NewPressureService() PressureService {
	return &service{}
}

func (*service) Create(pressure *domain.Pressure) (*domain.Pressure, error) {
	pressure.ID = uuid.New()
	return repo.Save(pressure)
}
