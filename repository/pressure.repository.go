package repository

import "api-pressure/domain"

type PressureRepository interface {
	Save(pressure *domain.Pressure) (*domain.Pressure, error)
	FindAll() ([]domain.Pressure, error)
}

type repo struct{}

var (
	pressures []domain.Pressure
)

func NewPressureRepository() PressureRepository {
	return &repo{}
}

func (*repo) Save(pressure *domain.Pressure) (*domain.Pressure, error) {
	println("Save in db (mock)")
	pressures = append(pressures, *pressure)
	return pressure, nil
}

func (*repo) FindAll() ([]domain.Pressure, error) {
	println("Read from db (mock)")
	return pressures, nil
}
