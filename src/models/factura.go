package models

import "time"

type EstadoFactura string

const (
	EstadoActivo  EstadoFactura = "ACTIVO"
	EstadoAnulado EstadoFactura = "ANULADO"
)

type Factura struct {
	ID                uint64        `gorm:"primaryKey;autoIncrement" json:"id"`
	CUF               string        `gorm:"type:varchar(100);unique;not null" json:"cuf"`
	CUFD              string        `gorm:"type:varchar(100);unique;not null" json:"cufd"`
	NitEmisor         string        `gorm:"type:varchar(100);not null" json:"nit_emisor"`
	CodigoSucursal    string        `gorm:"type:varchar(100);not null" json:"codigo_sucursal"`
	CodigoPuntoVenta  string        `gorm:"type:varchar(100);not null" json:"codigo_pv"`
	RazonSocialEmisor string        `gorm:"type:varchar(100);not null" json:"razon_social_emisor"`
	FechaEmision      time.Time     `gorm:"not null" json:"fecha_emision"`
	NitEmpresa        string        `gorm:"type:varchar(20);not null" json:"nit_empresa"`
	MontoTotal        float64       `gorm:"type:numeric(10,2);not null" json:"monto_total"`
	CodigoControl     string        `gorm:"type:varchar(50)" json:"codigo_control"`
	Estado            EstadoFactura `gorm:"type:estado_factura;not null" json:"estado"`
	CreatedAt         time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         *time.Time    `gorm:"index" json:"deleted_at"`

	// Relación 1:N con detalles
	Detalles []Detalle `gorm:"foreignKey:FacturaID" json:"detalles"`
}

// TableName devuelve el nombre correcto de la tabla en PostgreSQL
func (Factura) TableName() string {
	return "facturas"
}
