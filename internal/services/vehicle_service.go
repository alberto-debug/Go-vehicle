package services

import (
	"github.com/alberto-debug/Vehicle/api"
	"github.com/alberto-debug/Vehicle/internal/models"
	"github.com/alberto-debug/Vehicle/internal/storage"
)

// This function orchestrates the vehicle registration process.

func RegisterVehicle(input models.Vehicle) (models.Vehicle, error) {
	// Check CPF and Placa uniqueness
	if err := storage.ValidateVehicle(input.CPF, input.Placa); err != nil {
		return models.Vehicle{}, err
	}

	// Call external API
	apiData, err := api.GetVehicleData(input.Placa)
	if err != nil {
		return models.Vehicle{}, err
	}

	// merge the data
	input.Marca = apiData.Marca
	input.Modelo = apiData.Modelo
	input.Chassi = apiData.Chassi
	input.Licenciado = apiData.Licenciado
}
