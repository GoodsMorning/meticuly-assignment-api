package memory

import (
	"meticuly-assignment-api/internal/model"
)

type Memory struct {
	orders []model.ProductionOrder
	resources []model.Resource
	logs []model.Log
}

func New() *Memory {
	return &Memory{
		orders: mockOrders(),
		resources: mockResources(),
		logs: mockLogs(),
	}
}