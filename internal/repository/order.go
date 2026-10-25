package repository

import "meticuly-assignment-api/internal/model"

func (r *Repository) GetOrders() []model.ProductionOrder {
	return r.mockMemory.GetOrders()
}

func (r *Repository) CreateOrder(order model.ProductionOrder) model.ProductionOrder {
	return r.mockMemory.CreateOrder(order)
}

func (r *Repository) GetOrderSummary() []model.OrderSummary {
	counts := map[string]int{
		"queued":      0,
		"in-progress": 0,
		"blocked":     0,
		"finished":    0,
	}

	for _, order := range r.mockMemory.GetOrders() {
		counts[getOrderStatus(order)]++
	}

	return []model.OrderSummary{
		{Status: "queued", Count: counts["queued"]},
		{Status: "in-progress", Count: counts["in-progress"]},
		{Status: "blocked", Count: counts["blocked"]},
		{Status: "finished", Count: counts["finished"]},
	}
}

func getOrderStatus(order model.ProductionOrder) string {
	statuses := make([]string, 0, len(order.Operations))

	for _, operation := range order.Operations {
		statuses = append(statuses, operation.Status)
	}

	for _, status := range statuses {
		if status == "blocked" {
			return "blocked"
		}
	}

	allFinished := len(statuses) > 0
	for _, status := range statuses {
		if status != "finished" {
			allFinished = false
			break
		}
	}
	if allFinished {
		return "finished"
	}

	for _, status := range statuses {
		if status == "in-progress" {
			return "in-progress"
		}
	}

	for _, status := range statuses {
		if status == "finished" {
			return "in-progress"
		}
	}

	return "queued"
}