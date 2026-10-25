package service

import (
	"meticuly-assignment-api/internal/model"
)

func (s *Service) GetResources() []model.Resource {
	return s.repo.GetResouces()
}