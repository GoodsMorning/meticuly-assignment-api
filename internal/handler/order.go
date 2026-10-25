package handler

import (
	"meticuly-assignment-api/internal/model"
	"meticuly-assignment-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrders(c *gin.Context) {
	orders := h.service.GetOrder() 
	response.Success(c, 200, orders)
}

func (h *Handler) GetOrderSummary(c *gin.Context) {
	orderSummary := h.service.GetOrderSummary() 
	response.Success(c, 200, orderSummary)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var order model.ProductionOrder

	if err := c.ShouldBindJSON(&order); err != nil {
		response.Error(c, 400 ,err.Error(),)
		return
	}
	created := h.service.CreateOrder(order)

	response.Success(c, 201, created)
}