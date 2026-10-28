package handler

import (
	"meticuly-assignment-api/internal/model"
	"meticuly-assignment-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrders(c *gin.Context) {
	orders := h.service.GetOrders()
	response.Success(c, 200, orders)
}

func (h *Handler) GetOrderSummary(c *gin.Context) {
	orderSummary := h.service.GetOrderSummary()
	response.Success(c, 200, orderSummary)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var request model.CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	created, err := h.service.CreateOrder(request)
	if err != nil {
		response.Error(c,400, err.Error())
		return
	}

	response.Success(c, 201, created)
}
