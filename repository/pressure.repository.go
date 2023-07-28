package repository

import "api-pressure/domain"

type PressureRepository interface {
	Save(pressure *domain.Pressure) (*domain.Pressure, error)
	FindAll() ([]domain.Pressure, error)
}
