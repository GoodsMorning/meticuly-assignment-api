package repository


import "meticuly-assignment-api/internal/model"

func (r *Repository) GetMaterialStock() []model.MaterialStock {
	return r.mockMemory.GetMaterialStock()
}

func (r *Repository) UpdateMaterialStock(updated model.MaterialStock) error {
	return r.mockMemory.UpdateMaterialStock(updated)
}
