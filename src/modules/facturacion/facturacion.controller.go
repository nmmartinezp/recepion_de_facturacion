package facturacion

import (
	"app/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller maneja las rutas relacionadas con facturación
type Controller struct {
	service *Service
}

// =========================
// 📋 Obtener todas las facturas
// =========================

// GetFacturas obtiene todas las facturas registradas
// @Summary Lista de facturas
// @Description Obtiene todas las facturas con sus detalles
// @Tags Facturación
// @Produce json
// @Success 200 {array} models.Factura
// @Failure 500 {object} map[string]string
// @Router /facturacion/facturas [get]
// @Security BearerAuth
func (c *Controller) GetFacturas(ctx *gin.Context) {
	facturas, err := c.service.GetFacturas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las facturas"})
		return
	}
	ctx.JSON(http.StatusOK, facturas)
}

// =========================
// 🔍 Obtener una factura por CUF
// =========================

// GetFacturaByCUF obtiene una factura por su CUF
// @Summary Obtener factura por CUF
// @Description Devuelve una factura específica según su CUF
// @Tags Facturación
// @Produce json
// @Param cuf path string true "CUF de la factura"
// @Success 200 {object} models.Factura
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facturacion/facturas/{cuf} [get]
// @Security BearerAuth
func (c *Controller) GetFacturaByCUF(ctx *gin.Context) {
	cuf := ctx.Param("cuf")

	factura, err := c.service.GetFacturaByCUF(cuf)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, factura)
}

// =========================
// 🧾 Crear una nueva factura
// =========================

// CreateFactura crea una nueva factura con su detalle
// @Summary Registrar factura
// @Description Recibe una factura emitida y la almacena en la base de datos tras validación
// @Tags Facturación
// @Accept json
// @Produce json
// @Param factura body models.Factura true "Datos de la factura a registrar"
// @Success 201 {object} models.Factura
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facturacion/facturas [post]
// @Security BearerAuth
func (c *Controller) CreateFactura(ctx *gin.Context) {
	var input models.Factura

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	created, err := c.service.CreateFactura(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// =========================
// 🔧 Constructor del controlador
// =========================

func GetController() *Controller {
	service := GetService()
	controller := Controller{service: service}
	return &controller
}
