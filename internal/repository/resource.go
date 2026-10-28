package repository

import "meticuly-assignment-api/internal/model"

func (r *Repository) GetResouces() []model.Resource {
	return r.mockMemory.GetResources()
}

func (r *Repository) UpdateResource(updated model.Resource) error {
	return r.mockMemory.UpdateResource(updated)
}