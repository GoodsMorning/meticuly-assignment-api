package memory

import "meticuly-assignment-api/internal/model"

func (m *Memory) GetResources() []model.Resource {
	return m.resources
}

func mockResources() []model.Resource {
	return []model.Resource{
		{ID: "RES-001", Name: "Machine 01", Type: "Machine"},
		{ID: "RES-002", Name: "Assembly Station 01", Type: "Assembly Station"},
		{ID: "RES-003", Name: "QC Station 01", Type: "Quality Control"},
		{ID: "RES-004", Name: "Inspection Kit", Type: "Inspection Equipment"},
		{ID: "RES-005", Name: "Packaging Station 01", Type: "Packaging Station"},
	}
}
