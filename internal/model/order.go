package model

import "time"

type ProductionOrder struct {
	ID         string      `json:"id"`
	CreatedAt  time.Time   `json:"created_at"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Step      string      `json:"step"`
	Start     int         `json:"start"`
	Duration  int         `json:"duration"`
	Status    OrderStatus `json:"status"`
	Resources []string    `json:"resources"`
	Materials []Material  `json:"materials"`
}

type OrderSummary struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type CreateOrderRequest struct {
	ID         string     `json:"id"`
	Operations []StepType `json:"operations"`
}

type OrderStatus string

const (
	OrderStatusQueued     OrderStatus = "queued"
	OrderStatusInProgress OrderStatus = "in-progress"
	OrderStatusBlocked    OrderStatus = "blocked"
	OrderStatusFinished   OrderStatus = "finished"
)
