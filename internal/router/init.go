package router

import (
	"github.com/gin-gonic/gin"
	"meticuly-assignment-api/internal/handler"
)

func New(r *gin.Engine, h *handler.Handler){
	api := r.Group("/api")
	{
		api.GET("/health", h.Health)


		order := api.Group("/order")
		{
			order.GET("", h.GetOrders)
			order.GET("/summary", h.GetOrderSummary)
			
			order.POST("", h.CreateOrder)
		}
		// api.GET("/orders", h.GetAll)
	}
}