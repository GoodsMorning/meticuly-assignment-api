package service

import (
	"context"
	"meticuly-assignment-api/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(ctx context.Context) *Service {
	return &Service{
		repo: *repository.New(ctx),
	}
}