package service

import (
	"meticuly-assignment-api/internal/model"
)

func (s *Service) GetLogs() []model.Log {
	return s.repo.GetLogs()
}