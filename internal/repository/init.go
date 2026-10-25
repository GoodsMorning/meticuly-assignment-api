package repository

import (
	"context"
	"meticuly-assignment-api/internal/repository/memory"
)

type Repository struct {
	mockMemory memory.Memory
}

func New(ctx context.Context) *Repository {
	return &Repository{
		mockMemory: *memory.New(),
	}
}