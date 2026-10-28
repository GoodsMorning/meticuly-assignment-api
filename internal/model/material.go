package model

type Material struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type MaterialStock struct {
	Name     string  `json:"name"`
	Unit     string  `json:"unit"`
	Stock    float64 `json:"stock"`
	Consumed float64 `json:"consumed"`
	Reserved float64 `json:"reserved"`
}