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
	Repository repository.Repository
)

type service struct{}

func NewPressureService(repository repository.Repository) PressureService {
	Repository = repository
	return &service{}
}

func (*service) Create(pressure *domain.Pressure) (*domain.Pressure, error) {
	pressure.ID = uuid.New()
	return Repository.Save(pressure)
}

func (*service) FindAll() ([]domain.Pressure, error) {
	return Repository.FindAll()
}
