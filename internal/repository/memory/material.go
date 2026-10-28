package memory

import "meticuly-assignment-api/internal/model"


func (m *Memory) GetMaterialStock() []model.MaterialStock {
	return m.materials
}

func mockMaterial() []model.MaterialStock {
	return []model.MaterialStock{
		{Name: "Material A", Unit: "kg", Stock: 100, Consumed: 22, Reserved: 0},
		{Name: "Material B", Unit: "pcs", Stock: 60, Consumed: 8, Reserved: 6},
		{Name: "Fasteners", Unit: "pcs", Stock: 300, Consumed: 40, Reserved: 32},
		{Name: "Material C", Unit: "boxes", Stock: 40, Consumed: 0, Reserved: 4},
		{Name: "Calibration Fluid", Unit: "L", Stock: 5, Consumed: 0.5, Reserved: 0},
		{Name: "Label Roll", Unit: "roll", Stock: 6, Consumed: 0, Reserved: 1},
	}
}