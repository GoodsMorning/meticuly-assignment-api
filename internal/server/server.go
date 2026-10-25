package server

import (
	"context"
	"meticuly-assignment-api/internal/handler"
	"meticuly-assignment-api/internal/router"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	initializer.InitEnv()
// }

func RunServer() error {
	r := gin.Default()
	ctx := context.Background()

	serverHandler := handler.New(ctx)
	router.New(r, serverHandler)

	return r.Run()
}