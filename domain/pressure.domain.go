package domain

import "github.com/google/uuid"

type Pressure struct {
	ID        uuid.UUID `json:"id"`
	Diastolic uint8     `json:"diastolic"`
	Systolic  uint8     `json:"systolic"`
}
