package facturacion

import "time"

type FacturaRequest struct {
	CUF               string           `json:"cuf" validate:"required"`
	CUFD              string           `json:"cufd" validate:"required"`
	NitEmisor         string           `json:"nit_emisor" validate:"required"`
	CodigoSucursal    string           `json:"codigo_sucursal" validate:"required"`
	CodigoPuntoVenta  string           `json:"codigo_pv" validate:"required"`
	RazonSocialEmisor string           `json:"razon_social_emisor" validate:"required"`
	FechaEmision      time.Time        `json:"fecha_emision" validate:"required"`
	NitEmpresa        string           `json:"nit_empresa" validate:"required"`
	MontoTotal        float64          `json:"monto_total" validate:"required,gt=0"`
	CodigoControl     string           `json:"codigo_control"` // opcional
	Detalles          []DetalleRequest `json:"detalles" validate:"required,dive"`
}

type DetalleRequest struct {
	Descripcion    string  `json:"descripcion" validate:"required"`
	Cantidad       int     `json:"cantidad" validate:"required,gt=0"`
	PrecioUnitario float64 `json:"precio_unitario" validate:"required,gt=0"`
}
