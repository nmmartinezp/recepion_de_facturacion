package models

import (
	"time"

	"gorm.io/gorm"
)

type Factura struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CUF               string    `gorm:"size:50;not null;unique" json:"cuf"`
	NitEmisor         string    `gorm:"size:20;not null" json:"nitEmisor"`
	RazonSocialEmisor string    `gorm:"size:100;not null" json:"razonSocialEmisor"`
	FechaEmision      time.Time `json:"fechaEmision"`
	NitReceptor       string    `gorm:"size:20;not null" json:"nitReceptor"`
	MontoTotal        float64   `gorm:"type:decimal(10,2);not null" json:"montoTotal"`
	MontoDescuento    float64   `gorm:"type:decimal(10,2);default:0" json:"montoDescuento"`
	CodigoMetodoPago  int       `json:"codigoMetodoPago"`
	CodigoControl     string    `gorm:"size:50" json:"codigoControl"`
	FirmaDigital      string    `gorm:"type:text" json:"firmaDigital"`

	// Relación uno-a-muchos con Detalle
	Detalle []Detalle `gorm:"foreignKey:FacturaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"detalle"`

	// Campos de auditoría opcionales
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName devuelve el nombre correcto de la tabla en PostgreSQL
func (Factura) TableName() string {
	return "facturas"
}
