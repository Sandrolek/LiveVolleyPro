package main

import (
	"volley/internal/db"
	"volley/internal/router"
)

func main() {
	db.InitDB()

	r := router.SetupRouter(db.DB)

	r.Run(":8080")

	//docs.SwaggerInfo.BasePath = "/api/v1"
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
