package handler

import (
	"github.com/gin-gonic/gin"
	"meticuly-assignment-api/pkg/response"
)


func (h *Handler) GetResources(c *gin.Context) {
	orders := h.service.GetResources() 
	response.Success(c, 200, orders)
}