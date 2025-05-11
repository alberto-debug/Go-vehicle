package api

type ApiVehicle struct {
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Chassi     string `json:"chassi"`
	Licenciado bool   `json:"Licenciado"`
}
