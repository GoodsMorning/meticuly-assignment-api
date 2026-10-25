package router

import (
	"meticuly-assignment-api/internal/handler"

	"github.com/gin-gonic/gin"
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
		
		resource := api.Group("/resource")
		{
			resource.GET("", h.GetResources)
		}

		log := api.Group("/log")
		{
			log.GET("", h.GetLogs)
		}
	}
}