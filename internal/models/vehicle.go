package models

type Vehicle struct {
	ID           int    `json:"id"`
	Proprietario string `json:"proprietario"`
	CPF          string `json:"cpf"`
	Placa        string `json:"placa"`
	Marca        string `json:"marca"`
	Modelo       string `json:"modelo"`
	Chassi       string `json:"chassi"`
	Licenciado   bool   `json:"licenciado"`
}
