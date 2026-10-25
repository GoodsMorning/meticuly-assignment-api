package server

import (
	"context"
	"meticuly-assignment-api/internal/handler"
	"meticuly-assignment-api/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

// func init() {
// 	initializer.InitEnv()
// }

func RunServer() error {
	r := gin.Default()
	ctx := context.Background()

	r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))

	serverHandler := handler.New(ctx)
	router.New(r, serverHandler)

	return r.Run()
}