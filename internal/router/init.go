package router

import (
	"github.com/gin-gonic/gin"
	"meticuly-assignment-api/internal/handler"
)

func New(r *gin.Engine, h *handler.Handler){
	api := r.Group("/api")
	{
		api.GET("/health", h.Health)
	}
}