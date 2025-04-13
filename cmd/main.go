package main

import (
	"volley/internal/db"
)

func main() {
	db.InitDB()

	//r := router.SetupRouter(db.DB)
	//
	//docs.SwaggerInfo.BasePath = "/api/v1"
	//
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.Run(":8080")
}
