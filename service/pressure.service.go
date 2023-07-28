package service

import (
	"api-pressure/domain"
	"api-pressure/repository"
	"errors"

	"github.com/google/uuid"
)

type PressureService interface {
	Validate(post *domain.Pressure) error
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

func (*service) Validate(pressure *domain.Pressure) error {
	if pressure == nil {
		err := errors.New("The pressure is empty")
		return err
	}
	if pressure.Diastolic == 0 || pressure.Systolic == 0 {
		err := errors.New("Values are in error")
		return err
	}
	return nil
}

func (*service) Create(pressure *domain.Pressure) (*domain.Pressure, error) {
	pressure.ID = uuid.New()
	return Repository.Save(pressure)
}

func (*service) FindAll() ([]domain.Pressure, error) {
	return Repository.FindAll()
}
