package constant

import "meticuly-assignment-api/internal/model"

var StepCatalog = map[model.StepType]model.StepTemplate{
	model.StepMachining1: {
		ID:          model.StepMachining1,
		Name:        "Machining-1",
		Description: "Machine and prepare raw components.",
		Duration:    90,
		Resources:   []string{"Machine 01"},
		Materials: []model.MaterialRequirement{
			{
				Name:        "Material A",
				Unit:        "kg",
				Numerator:   12,
				Denominator: 100,
			},
		},
	},

	model.StepMachining2: {
		ID:          model.StepMachining2,
		Name:        "Machining-2",
		Description: "Machine and prepare raw components.",
		Duration:    90,
		Resources:   []string{"Machine 02"},
		Materials: []model.MaterialRequirement{
			{
				Name:        "Material B",
				Unit:        "kg",
				Numerator:   12,
				Denominator: 100,
			},
		},
	},

	model.StepAssembly: {
		ID:          model.StepAssembly,
		Name:        "Assembly",
		Description: "Assemble components into the finished product.",
		Duration:    90,
		Resources:   []string{"Assembly Station 01"},
		Materials: []model.MaterialRequirement{
			{
				Name:        "Material B",
				Unit:        "pcs",
				Numerator:   8,
				Denominator: 100,
			},
			{
				Name:        "Fasteners",
				Unit:        "pcs",
				Numerator:   40,
				Denominator: 100,
			},
		},
	},

	model.StepQualityCheck: {
		ID:          model.StepQualityCheck,
		Name:        "Quality Check",
		Description: "Inspect product quality and specifications.",
		Duration:    60,
		Resources:   []string{"QC Station 01", "Inspection Kit"},
		Materials: []model.MaterialRequirement{
			{
				Name:        "Calibration Fluid",
				Unit:        "L",
				Numerator:   5,
				Denominator: 1000,
			},
		},
	},

	model.StepPackaging: {
		ID:          model.StepPackaging,
		Name:        "Packaging",
		Description: "Pack finished products for delivery.",
		Duration:    60,
		Resources:   []string{"Packaging Station 01"},
		Materials: []model.MaterialRequirement{
			{
				Name:        "Material C",
				Unit:        "boxes",
				Numerator:   1,
				Denominator: 1,
			},
			{
				Name:        "Label Roll",
				Unit:        "roll",
				Numerator:   1,
				Denominator: 100,
			},
		},
	},
}