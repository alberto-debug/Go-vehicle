package handlers

import (
	"net/http"

	"github.com/alberto-debug/Vehicle/internal/models"
)

// This is the HTTP handler for the /veiculos route. It uses the service layer to register the vehicle.

func CreateVehicle(w http.ResponseWriter, r *http.ResponseWriter) {
	var input models.Vehicle
}
