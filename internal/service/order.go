package service

import (
	"meticuly-assignment-api/internal/model"
)

func (s *Service) GetOrder() []model.ProductionOrder {
	return s.repo.GetOrder()
}

func (s *Service) GetOrderSummary() []model.OrderSummary {
	return s.repo.GetOrderSummary()
}

func (s *Service) CreateOrder(order model.ProductionOrder) model.ProductionOrder {
	return s.repo.CreateOrder(order)
}