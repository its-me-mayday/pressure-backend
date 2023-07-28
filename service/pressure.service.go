package service

import (
	"api-pressure/domain"
	"api-pressure/repository"

	"github.com/google/uuid"
)

type PressureService interface {
	Create(pressure *domain.Pressure) (*domain.Pressure, error)
	FindAll() ([]domain.Pressure, error)
}

var (
	pressureRepository repository.PressureRepository
)

type service struct{}

func NewPressureService(repository repository.PressureRepository) PressureService {
	pressureRepository = repository
	return &service{}
}

func (*service) Create(pressure *domain.Pressure) (*domain.Pressure, error) {
	pressure.ID = uuid.New()
	return pressureRepository.Save(pressure)
}

func (*service) FindAll() ([]domain.Pressure, error) {
	return pressureRepository.FindAll()
}
