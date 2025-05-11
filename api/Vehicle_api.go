package api

import (
	"fmt"
	"net/http"
)

type ApiVehicle struct {
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Chassi     string `json:"chassi"`
	Licenciado bool   `json:"Licenciado"`
}

func GetVehicleData(placa string) (ApiVehicle, error) {
	url := fmt.Sprint("https://my.api.mockaroo.com/veiculos?key=55ad1cd0&placa=%s", placa)

	resp, err := http.Get(url)
	if err != nil {
		return ApiVehicle{}, err
	}

	defer resp.Body.Close()
}
