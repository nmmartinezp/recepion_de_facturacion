package main

import (
	"app/src/modules/facturacion"

	_ "app/src/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupApp() *gin.Engine {
	app := gin.Default()

	router := app.Group("/api/v1")

	facturacion.Router(router)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
