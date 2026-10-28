package memory

import (
	"fmt"
	"meticuly-assignment-api/internal/model"
)

func (m *Memory) GetResources() []model.Resource {
	return m.resources
}

func (m *Memory) UpdateResource(updated model.Resource) error {
	for i, resource := range m.resources {
		if resource.ID == updated.ID {
			m.resources[i] = updated
			return nil
		}
	}

	return fmt.Errorf("resource not found: %s", updated.ID)
}

func mockResources() []model.Resource {
	return []model.Resource{
		{
			ID:          "RES-001",
			Name:        "Machine 01",
			Status:      model.ResourceStatusBusy,
			Utilization: 85.5,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Shared machining resource",
		},
		{
			ID:          "RES-006",
			Name:        "Machine 02",
			Status:      model.ResourceStatusBusy,
			Utilization: 70.5,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Shared machining resource",
		},
		{
			ID:          "RES-002",
			Name:        "Assembly Station 01",
			Status:      model.ResourceStatusBusy,
			Utilization: 65.0,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Processing multiple assembly orders",
		},
		{
			ID:          "RES-003",
			Name:        "QC Station 01",
			Status:      model.ResourceStatusBusy,
			Utilization: 72.0,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Shared quality control resource",
		},
		{
			ID:          "RES-004",
			Name:        "Inspection Kit",
			Status:      model.ResourceStatusBusy,
			Utilization: 50.0,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Shared inspection equipment",
		},
		{
			ID:          "RES-005",
			Name:        "Packaging Station 01",
			Status:      model.ResourceStatusBusy,
			Utilization: 45.0,
			Orders:      []string{"ORD-001", "ORD-002"},
			Note:        "Assigned to two packaging orders",
		},
	}
}
