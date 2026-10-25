package model

type ProductionOrder struct {
	ID         string      `json:"id"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Step      string     `json:"step"`
	Start     int        `json:"start"`
	Duration  int        `json:"duration"`
	Status    string     `json:"status"`
	Resources []string   `json:"resources"`
	Materials []Material `json:"materials"`
}

type OrderSummary struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}