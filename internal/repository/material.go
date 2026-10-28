package repository


import "meticuly-assignment-api/internal/model"

func (r *Repository) GetMaterialStock() []model.MaterialStock {
	return r.mockMemory.GetMaterialStock()
}