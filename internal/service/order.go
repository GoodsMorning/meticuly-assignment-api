package service

import (
	"fmt"
	"meticuly-assignment-api/internal/model"
	"sort"
	"time"
)

func (s *Service) GetOrders() []model.ProductionOrder {
	orders := s.repo.GetOrders()

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.After(orders[j].CreatedAt)
	})

	return orders

}

func (s *Service) GetOrderSummary() []model.OrderSummary {
	return s.repo.GetOrderSummary()
}

func (s *Service) CreateOrder(order model.ProductionOrder) model.ProductionOrder {
	created := s.repo.CreateOrder(order)
	s.repo.WriteLog(model.Log{
		OrderID: created.ID,
		Message: fmt.Sprintf("Production order %s created", created.ID),
		Level:   "info",
		Time:    time.Now().Format("15:04:05"),
	})
	return created
}