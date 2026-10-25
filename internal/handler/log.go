package handler

import (
	"github.com/gin-gonic/gin"
	"meticuly-assignment-api/pkg/response"
)

func (h *Handler) GetLogs(c *gin.Context) {
	logs := h.service.GetLogs() 
	response.Success(c, 200, logs)
}