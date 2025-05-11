package storage

import (
	"errors"

	"github.com/alberto-debug/Vehicle/internal/models"
)

// Basic in memory storage
// Create a variable named vehicleDB that is an empty list (slice) of Vehicle structs from the models package

var vehicleDB = []models.Vehicle{}

func ValidateVehicle(cpf, placa string) error {
	for _, v := range vehicleDB {
		if v.CPF == cpf {
			return errors.New("CPF already exists")
		}
		if v.Placa == placa {
			return errors.New("placa already exists")
		}
	}

	return nil
}

func SaveVehicle(vehicle models.Vehicle) (models.Vehicle, error) {
	vehicle.ID = len(vehicleDB) + 1
	vehicleDB = append(vehicleDB, vehicle)

	return vehicle, nil
}
