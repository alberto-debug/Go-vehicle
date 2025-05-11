package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alberto-debug/Vehicle/internal/models"
	"github.com/alberto-debug/Vehicle/internal/services"
)

// This is the HTTP handler for the /veiculos route. It uses the service layer to register the vehicle.

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var input models.Vehicle

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	vehicle, err := services.RegisterVehicle(input)
	if err != nil {

		http.Error(w, err.Error(), http.StatusConflict)
		return

	}
}
