package memory

import (
	"fmt"
	"meticuly-assignment-api/internal/constant"
	"meticuly-assignment-api/internal/model"
	"strings"
	"time"
)

func (m *Memory) GetOrders() []model.ProductionOrder {
	return m.orders
}

func (m *Memory) CreateOrder(
	order model.CreateOrderRequest,
) (model.ProductionOrder, error) {
	baseID := strings.TrimSpace(order.ID)

	if baseID == "" {
		baseID = fmt.Sprintf("ORD-%03d", len(m.orders)+1)
	}

	candidate := baseID
	for suffix := 1; m.orderIDExists(candidate); suffix++ {
		candidate = fmt.Sprintf("%s(%d)", baseID, suffix)
	}
	// Convert requested step types into production operations.
	operations := make([]model.Operation, 0, len(order.Operations))

	start := 0
	for _, stepType := range order.Operations {
		template, ok := constant.StepCatalog[stepType]
		if !ok {
			return model.ProductionOrder{}, fmt.Errorf(
				"unsupported operation: %s", stepType,
			)
		}

		operations = append(operations, model.Operation{
			Step:      string(template.ID),
			Start:     start,
			Duration:  template.Duration,
			Status:    model.OrderStatusQueued,
			Resources: append([]string(nil), template.Resources...),
			Materials: convertMaterials(template.Materials),
		})

		start += template.Duration
	}

	created := model.ProductionOrder{
		ID:         candidate,
		CreatedAt:  time.Now(),
		Operations: operations,
	}

	m.orders = append(m.orders, created)

	return created, nil
}

func (m *Memory) orderIDExists(id string) bool {
	for _, order := range m.orders {
		if order.ID == id {
			return true
		}
	}
	return false
}

func convertMaterials(
	materials []model.MaterialRequirement,
) []model.Material {
	result := make([]model.Material, 0, len(materials))

	for _, material := range materials {
		quantity := float64(material.Numerator) /
			float64(material.Denominator)

		result = append(result, model.Material{
			Name:     material.Name,
			Unit:     material.Unit,
			Quantity: quantity,
		})
	}

	return result
}

func mockOrders() []model.ProductionOrder {
	return []model.ProductionOrder{
		{
			ID:        "ORD-001",
			CreatedAt: time.Date(2026, 10, 24, 9, 30, 0, 0, time.Local),
			Operations: []model.Operation{
				{
					Step:      "Machining-1",
					Start:     -270,
					Duration:  90,
					Status:    "finished",
					Resources: []string{"Machine 01"},
					Materials: []model.Material{
						{Name: "Material A", Quantity: 12, Unit: "kg"},
					},
				},
				{
					Step:      "Assembly",
					Start:     -180,
					Duration:  90,
					Status:    "finished",
					Resources: []string{"Assembly Station 01"},
					Materials: []model.Material{
						{Name: "Material B", Quantity: 8, Unit: "pcs"},
						{Name: "Fasteners", Quantity: 40, Unit: "pcs"},
					},
				},
				{
					Step:      "Quality Check",
					Start:     -90,
					Duration:  60,
					Status:    "finished",
					Resources: []string{"QC Station 01", "Inspection Kit"},
					Materials: []model.Material{
						{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"},
					},
				},
				{
					Step:      "Packaging",
					Start:     -30,
					Duration:  60,
					Status:    "in-progress",
					Resources: []string{"Packaging Station 01"},
					Materials: []model.Material{
						{Name: "Material C", Quantity: 4, Unit: "boxes"},
						{Name: "Label Roll", Quantity: 1, Unit: "roll"},
					},
				},
			},
		},
		{
			ID:        "ORD-002",
			CreatedAt: time.Date(2026, 10, 25, 14, 15, 0, 0, time.Local),
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: -150, Duration: 90, Status: "finished",
					Resources: []string{"Machine 01"},
					Materials: []model.Material{{
						Name:     "Material A",
						Quantity: 10,
						Unit:     "kg"}},
				},
				{
					Step:  "Assembly",
					Start: -60, Duration: 120, Status: "in-progress",
					Resources: []string{"Assembly Station 01"},
					Materials: []model.Material{{
						Name:     "Material B",
						Quantity: 6, Unit: "pcs"}, {
						Name:     "Fasteners",
						Quantity: 32,
						Unit:     "pcs"}},
				},
				{
					Step:  "Quality Check",
					Start: 60, Duration: 60, Status: "queued",
					Resources: []string{"QC Station 01", "Inspection Kit"},
					Materials: []model.Material{{
						Name:     "Calibration Fluid",
						Quantity: 0.5,
						Unit:     "L"}},
				},
				{
					Step:  "Packaging",
					Start: 120, Duration: 60, Status: "queued",
					Resources: []string{"Packaging Station 01"},
					Materials: []model.Material{
						{
							Name:     "Material C",
							Quantity: 3, Unit: "boxes"},
						{
							Name:     "Label Roll",
							Quantity: 1,
							Unit:     "roll",
						}},
				},
			},
		},
	}
}
