package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiVehicle struct {
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Chassi     string `json:"chassi"`
	Licenciado bool   `json:"Licenciado"`
}

// indicates that the function GetVehicleData takes a placa (a vehicle's license plate) as input, and it returns two things
func GetVehicleData(placa string) (ApiVehicle, error) {
	url := fmt.Sprint("https://my.api.mockaroo.com/veiculos?key=55ad1cd0&placa=%s", placa)

	// This line sends an HTTP GET request to the constructed URL. The response is
	// stored in resp, and any error that occurs during the request is stored in err
	resp, err := http.Get(url)
	// If an error occurs while making the GET request (e.g.,
	// network issue, invalid URL), the function returns an empty ApiVehicle struct and the error
	if err != nil {
		return ApiVehicle{}, err
	}

	// close the request
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ApiVehicle{}, fmt.Errorf("Api error: %s", resp.Status)
	}

	// This variable will be used to store the vehicle data retrieved from the API response
	var vehicle ApiVehicle
	if err := json.NewDecoder(resp.Body).Decode(&vehicle); err != nil {
		return ApiVehicle{}, err
	}

	return vehicle, nil
}
