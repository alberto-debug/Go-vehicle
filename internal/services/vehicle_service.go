package services

import (
	"github.com/alberto-debug/Vehicle/internal/models"
	"github.com/alberto-debug/Vehicle/internal/storage"
)

// This function orchestrates the vehicle registration process.

func RegisterVehicle(input models.Vehicle) (models.Vehicle, error) {
	if err := storage.ValidateVehicle(input.CPF, input.Placa); err != nil {
		return models.Vehicle{}, err
	}
}
