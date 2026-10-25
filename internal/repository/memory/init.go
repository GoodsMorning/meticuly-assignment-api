package memory

import (
	"meticuly-assignment-api/internal/model"
)

type Memory struct {
	orders []model.ProductionOrder
}

func New() *Memory {
	return &Memory{
		orders: mockOrders(),
	}
}