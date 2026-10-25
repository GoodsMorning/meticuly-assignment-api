package repository

import "meticuly-assignment-api/internal/model"

func (r *Repository) GetResouces() []model.Resource {
	return r.mockMemory.GetResources()
}