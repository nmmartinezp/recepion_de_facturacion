package facturacion

import (
	"app/src/middlewares"

	"github.com/gin-gonic/gin"
)

// Router configura las rutas del módulo de facturación
func Router(router *gin.RouterGroup) {
	controller := GetController()

	r := router.Group("/facturacion")
	r.Use(middlewares.AuthMiddleware())
	{
		// 📋 Obtener todas las facturas
		r.GET("/facturas", controller.GetFacturas)

		// 🔍 Obtener una factura por CUF
		r.GET("/facturas/:cuf", controller.GetFacturaByCUF)

		// 🧾 Registrar una nueva factura
		r.POST("/facturas", controller.CreateFactura)
	}
}
