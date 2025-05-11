package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alberto-debug/Vehicle/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/veiculos", handlers.CreateVehicle).Methods("POST")

	log.Println("Server in running on http://8080")
	if err := http.ListenAndServe(":8080"); err != nil {
		log.Fatal(err)
	}
}
