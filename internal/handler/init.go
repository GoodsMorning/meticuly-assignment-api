package handler

import (
	"context"
	"meticuly-assignment-api/internal/service"
)

type Handler struct {
	service	service.Service
}


func New(ctx context.Context) *Handler {
	return &Handler{
		service: *service.New(ctx),
	}
}