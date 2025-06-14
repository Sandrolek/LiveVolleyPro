package main

import (
	"volley/cmd/docs"
	"volley/internal/db"
	"volley/internal/router"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.InitDB()

	r := gin.Default()

	router.SetupRouter(r, db.DB)

	r.Run(":8080")

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
