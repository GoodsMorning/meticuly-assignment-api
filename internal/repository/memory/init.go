package memory

import (
	"meticuly-assignment-api/internal/model"
)

type Memory struct {
	orders []model.ProductionOrder
	resources []model.Resource
	materials []model.MaterialStock
	logs []model.Log
}

func New() *Memory {
	return &Memory{
		orders: mockOrders(),
		resources: mockResources(),
		materials: mockMaterial(),
		logs: mockLogs(),
	}
}