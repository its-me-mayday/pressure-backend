package repository

import "api-pressure/domain"

type mockRepository struct{}

var (
	pressures []domain.Pressure
)

func NewPressureRepository() PressureRepository {
	return &mockRepository{}
}

func (*mockRepository) Save(pressure *domain.Pressure) (*domain.Pressure, error) {
	println("Save in db (mock)")
	pressures = append(pressures, *pressure)
	return pressure, nil
}

func (*mockRepository) FindAll() ([]domain.Pressure, error) {
	println("Read from db (mock)")
	return pressures, nil
}
