package handler

import (
	"github.com/gin-gonic/gin"
	"meticuly-assignment-api/pkg/response"
)


func (h *Handler) GetStock(c *gin.Context) {
	stock := h.service.GetStock() 
	response.Success(c, 200, stock)
}