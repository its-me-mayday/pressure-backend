package service

import (
	"api-pressure/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPressure(t *testing.T) {
	testService := NewPressureService(nil)
	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The pressure is empty", err.Error())
}

func TestValidateZeroValue(t *testing.T) {
	pressure := domain.Pressure{ID: uuid.New(), Diastolic: 0, Systolic: 1}
	testService := NewPressureService(nil)
	err := testService.Validate(&pressure)

	assert.NotNil(t, err)
	assert.Equal(t, "Values are in error", err.Error())
}
