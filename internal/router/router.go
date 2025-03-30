package router

import (
	"volley/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Пример: эндпоинты для users
		api.GET("/users", handlers.GetAllUsers)
		api.POST("/users", handlers.CreateUser)
	}

	return r
}
