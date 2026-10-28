package service

import (
	"meticuly-assignment-api/internal/model"
)

func (s *Service) GetStock() model.Stock {
	resources := s.repo.GetResouces()
	materialStocks := s.repo.GetMaterialStock()

	return model.Stock{
		Resources: resources,
		Materials: materialStocks,
	}
}