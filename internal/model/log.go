package model

type Log struct {
	Time    string `json:"time"`
	OrderID string `json:"orderId"`
	Message string `json:"message"`
	Level   string `json:"level"`
}