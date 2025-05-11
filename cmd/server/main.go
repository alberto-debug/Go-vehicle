package main

import (
	"github.com/gorilla/mux"

	"github.com/alberto-debug/Vehicle/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/veiculos", handlers.CreateVehicle).Methods("POST")
}
