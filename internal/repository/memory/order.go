package memory

import (
	"fmt"
	"meticuly-assignment-api/internal/model"
)

func (m *Memory) GetOrders() []model.ProductionOrder {
	return m.orders
}

func (m *Memory) CreateOrder(order model.ProductionOrder) model.ProductionOrder {
	order.ID = fmt.Sprintf("ORD-%03d", len(m.orders)+1)
	m.orders = append(m.orders, order)

	return order
}

func mockOrders() []model.ProductionOrder {
	return []model.ProductionOrder{
		{
			ID: "ORD-001",
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
			ID: "ORD-002",
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: -150, Duration: 90, Status: "finished", Resources: []string{"Machine 01"}, Materials: []model.Material{{Name: "Material A", Quantity: 10, Unit: "kg"}}},
				{
					Step:  "Assembly",
					Start: -60, Duration: 120, Status: "in-progress", Resources: []string{"Assembly Station 01"}, Materials: []model.Material{{Name: "Material B", Quantity: 6, Unit: "pcs"}, {Name: "Fasteners", Quantity: 32, Unit: "pcs"}}},
				{
					Step:  "Quality Check",
					Start: 60, Duration: 60, Status: "queued", Resources: []string{"QC Station 01", "Inspection Kit"}, Materials: []model.Material{{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"}}},
				{
					Step:  "Packaging",
					Start: 120, Duration: 60, Status: "queued", Resources: []string{"Packaging Station 01"}, Materials: []model.Material{{Name: "Material C", Quantity: 3, Unit: "boxes"}, {Name: "Label Roll", Quantity: 1, Unit: "roll"}}},
			},
		},
		{
			ID: "ORD-003",
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: -120, Duration: 90, Status: "finished", Resources: []string{"Machine 02"}, Materials: []model.Material{{Name: "Material A", Quantity: 15, Unit: "kg"}}},
				{
					Step:  "Assembly",
					Start: -30, Duration: 75, Status: "blocked", Resources: []string{"Assembly Station 02"}, Materials: []model.Material{{Name: "Material B", Quantity: 10, Unit: "pcs"}, {Name: "Fasteners", Quantity: 50, Unit: "pcs"}}},
				{
					Step:  "Quality Check",
					Start: 45, Duration: 60, Status: "queued", Resources: []string{"QC Station 01", "Inspection Kit"}, Materials: []model.Material{{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"}}},
				{
					Step:  "Packaging",
					Start: 105, Duration: 60, Status: "queued", Resources: []string{"Packaging Station 02"}, Materials: []model.Material{{Name: "Material C", Quantity: 5, Unit: "boxes"}}},
			},
		},
		{
			ID: "ORD-004",
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: 10, Duration: 150, Status: "queued", Resources: []string{"Machine 01"}, Materials: []model.Material{{Name: "Material B", Quantity: 20, Unit: "kg"}}},
				{
					Step:  "Assembly",
					Start: 160, Duration: 90, Status: "queued", Resources: []string{"Assembly Station 01"}, Materials: []model.Material{{Name: "Material B", Quantity: 12, Unit: "pcs"}, {Name: "Fasteners", Quantity: 60, Unit: "pcs"}}},
				{
					Step:  "Quality Check",
					Start: 250, Duration: 45, Status: "queued", Resources: []string{"QC Station 01", "Inspection Kit"}, Materials: []model.Material{{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"}}},
				{
					Step:  "Packaging",
					Start: 295, Duration: 60, Status: "queued", Resources: []string{"Packaging Station 01"}, Materials: []model.Material{{Name: "Material C", Quantity: 6, Unit: "boxes"}, {Name: "Label Roll", Quantity: 1, Unit: "roll"}}},
			},
		},
		{
			ID: "ORD-005",
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: 90, Duration: 90, Status: "queued", Resources: []string{"Machine 02"}, Materials: []model.Material{{Name: "Material A", Quantity: 14, Unit: "kg"}}},
				{
					Step:  "Assembly",
					Start: 180, Duration: 90, Status: "queued", Resources: []string{"Assembly Station 02"}, Materials: []model.Material{{Name: "Material B", Quantity: 9, Unit: "pcs"}, {Name: "Fasteners", Quantity: 45, Unit: "pcs"}}},
				{
					Step:  "Quality Check",
					Start: 270, Duration: 45, Status: "queued", Resources: []string{"QC Station 01", "Inspection Kit"}, Materials: []model.Material{{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"}}},
				{
					Step:  "Packaging",
					Start: 315, Duration: 45, Status: "queued", Resources: []string{"Packaging Station 02"}, Materials: []model.Material{{Name: "Material C", Quantity: 4, Unit: "boxes"}}},
			},
		},
		{
			ID: "ORD-006",
			Operations: []model.Operation{
				{
					Step:  "Machining-1",
					Start: -90, Duration: 90, Status: "finished", Resources: []string{"Machine 02"}, Materials: []model.Material{{Name: "Material A", Quantity: 14, Unit: "kg"}}},
				{
					Step:  "Machining-2",
					Start: 0, Duration: 90, Status: "queued", Resources: []string{"Machine 02"}, Materials: []model.Material{{Name: "Material A", Quantity: 14, Unit: "kg"}}},
				{
					Step:  "Machining-1",
					Start: 0, Duration: 90, Status: "queued", Resources: []string{"Machine 02"}, Materials: []model.Material{{Name: "Material A", Quantity: 14, Unit: "kg"}}},
				{
					Step:  "Assembly",
					Start: 180, Duration: 90, Status: "queued", Resources: []string{"Assembly Station 02"}, Materials: []model.Material{{Name: "Material B", Quantity: 9, Unit: "pcs"}, {Name: "Fasteners", Quantity: 45, Unit: "pcs"}}},
				{
					Step:  "Quality Check",
					Start: 270, Duration: 45, Status: "queued", Resources: []string{"QC Station 01", "Inspection Kit"}, Materials: []model.Material{{Name: "Calibration Fluid", Quantity: 0.5, Unit: "L"}}},
				{
					Step:  "Packaging",
					Start: 315, Duration: 45, Status: "queued", Resources: []string{"Packaging Station 02"}, Materials: []model.Material{{Name: "Material C", Quantity: 4, Unit: "boxes"}}},
			},
		},
	}
}
