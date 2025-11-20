package facturacion

import (
	cn "app/src/databases"
	"app/src/models"
	"context"
	"errors"
	"fmt"

	"app/src/modules/common"

	"gorm.io/gorm"

	punto_venta "app/src/proto/punto_venta"

	"google.golang.org/grpc"
)

type Service struct {
	db       *gorm.DB
	GRPCConn punto_venta.PuntoVentaServiceClient
}

// ===========================
// VALIDACIÓN Y REGISTRO DE FACTURAS
// ===========================

// CreateFactura recibe una factura con su detalle, valida campos esenciales y la registra en la base de datos
func (s *Service) CreateFactura(req FacturaRequest) (models.Factura, error) {
	var validationErrors []string

	// =============================
	// VALIDACIÓN DE CUFD VÍA gRPC
	// =============================
	valido, motivo, err := s.ValidarCufd(req.NitEmisor, req.CodigoSucursal, req.CodigoPuntoVenta, req.CUFD)
	if err != nil {
		validationErrors = append(validationErrors, "Error al validar CUFD por gRPC: "+err.Error())
	} else if !valido {
		validationErrors = append(validationErrors, "CUFD inválido: "+motivo)
	}

	// =============================
	// VALIDACIONES DE CAMPOS VACÍOS
	// =============================
	if req.CUF == "" {
		validationErrors = append(validationErrors, "El campo CUF es obligatorio y está vacío")
	}
	if req.CUFD == "" {
		validationErrors = append(validationErrors, "El campo CUFD es obligatorio y está vacío")
	}
	if req.NitEmisor == "" {
		validationErrors = append(validationErrors, "El campo Nit Emisor es obligatorio y está vacío")
	}
	if req.CodigoSucursal == "" {
		validationErrors = append(validationErrors, "El campo Código Sucursal es obligatorio y está vacío")
	}
	if req.CodigoPuntoVenta == "" {
		validationErrors = append(validationErrors, "El campo Código Punto de Venta es obligatorio y está vacío")
	}
	if req.RazonSocialEmisor == "" {
		validationErrors = append(validationErrors, "El campo Razón Social Emisor es obligatorio y está vacío")
	}
	if req.NitEmpresa == "" {
		validationErrors = append(validationErrors, "El campo NIT Empresa es obligatorio y está vacío")
	}
	if req.MontoTotal <= 0 {
		validationErrors = append(validationErrors, "El Monto Total debe ser mayor a 0")
	}

	if len(req.Detalles) == 0 {
		validationErrors = append(validationErrors, "La factura debe contener al menos un detalle")
	}

	// =============================
	// VALIDACIÓN DE DETALLES
	// =============================
	var totalCalculado float64 = 0

	for i, d := range req.Detalles {

		if d.Descripcion == "" {
			validationErrors = append(validationErrors,
				fmt.Sprintf("El detalle %d tiene la descripción vacía", i+1))
		}

		if d.Cantidad <= 0 {
			validationErrors = append(validationErrors,
				fmt.Sprintf("El detalle %d tiene cantidad inválida (debe ser mayor a 0)", i+1))
		}

		if d.PrecioUnitario <= 0 {
			validationErrors = append(validationErrors,
				fmt.Sprintf("El detalle %d tiene precio unitario inválido (debe ser mayor a 0)", i+1))
		}

		// Sumar al total
		totalCalculado += float64(d.Cantidad) * d.PrecioUnitario
	}

	// =============================
	// VALIDAR QUE EL MONTO TOTAL COINCIDA
	// =============================
	if totalCalculado != req.MontoTotal {
		validationErrors = append(validationErrors,
			fmt.Sprintf("El monto total declarado (%.2f) no coincide con el total calculado (%.2f)",
				req.MontoTotal, totalCalculado))
	}

	// =============================
	// DETERMINAR ESTADO
	// =============================
	estado := models.EstadoActivo
	if len(validationErrors) > 0 {
		estado = models.EstadoAnulado
	}

	// =============================
	// MAPEO A MODELO
	// =============================
	factura := models.Factura{
		CUF:               req.CUF,
		CUFD:              req.CUFD,
		NitEmisor:         req.NitEmisor,
		CodigoSucursal:    req.CodigoSucursal,
		CodigoPuntoVenta:  req.CodigoPuntoVenta,
		RazonSocialEmisor: req.RazonSocialEmisor,
		FechaEmision:      req.FechaEmision,
		NitEmpresa:        req.NitEmpresa,
		MontoTotal:        req.MontoTotal,
		CodigoControl:     req.CodigoControl,
		Estado:            estado,
	}

	// Añadir detalles al modelo
	for _, d := range req.Detalles {
		factura.Detalles = append(factura.Detalles, models.Detalle{
			Descripcion:    d.Descripcion,
			Cantidad:       d.Cantidad,
			PrecioUnitario: d.PrecioUnitario,
		})
	}

	// =============================
	// GUARDAR EN BASE DE DATOS
	// =============================
	if err := s.db.Create(&factura).Error; err != nil {
		return models.Factura{}, err
	}

	// =============================
	// ENVIAR MENSAJE A RABBITMQ
	// =============================
	message := common.BillingValidationMessage{
		EventID:   fmt.Sprintf("%d", factura.ID),
		EmpresaID: factura.NitEmpresa,
		NIT:       factura.NitEmisor,
		Factura:   factura,
		Errores:   validationErrors,
		Resultado: func() string {
			if estado == models.EstadoActivo {
				return "VALID"
			}
			return "INVALID"
		}(),
	}

	if err := common.PublishValidationResult(message); err != nil {
		fmt.Println("Error enviando mensaje de RabbitMQ:", err)
	}
	return factura, nil
}

// ===========================
// CONSULTAS DE FACTURAS
// ===========================

func (s *Service) GetFacturas() ([]models.Factura, error) {
	var facturas []models.Factura
	err := s.db.Preload("Detalle").Find(&facturas).Error
	if err != nil {
		return nil, err
	}
	return facturas, nil
}

func (s *Service) GetFacturaByCUF(cuf string) (models.Factura, error) {
	var factura models.Factura
	err := s.db.Preload("Detalle").Where("cuf = ?", cuf).First(&factura).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Factura{}, fmt.Errorf("no se encontró la factura con CUF %s", cuf)
		}
		return models.Factura{}, err
	}
	return factura, nil
}

func (s *Service) ValidarCufd(nit, sucursal, puntoVenta, cufd string) (bool, string, error) {

	req := &punto_venta.ValidarCufdRequest{
		Nit:              nit,
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: puntoVenta,
		Cufd:             cufd,
	}

	resp, err := s.GRPCConn.ValidarCufd(context.Background(), req)
	if err != nil {
		return false, "", err
	}

	return resp.Valido, resp.Motivo, nil
}

// ===========================
// CONSTRUCTOR DEL SERVICIO
// ===========================

func GetService() *Service {
	conn, err := grpc.NewClient("servicio-empresas-cufd:50053")
	if err != nil {
		panic("No se pudo conectar al servicio PuntoVenta: " + err.Error())
	}

	client := punto_venta.NewPuntoVentaServiceClient(conn)
	db := cn.DBPOSTGRES()
	return &Service{db: db, GRPCConn: client}
}
