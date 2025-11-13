package models

type Detalle struct {
	ID             uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FacturaID      uint    `gorm:"not null;index" json:"facturaId"`
	Descripcion    string  `gorm:"size:200;not null" json:"descripcion"`
	Cantidad       int     `gorm:"not null" json:"cantidad"`
	PrecioUnitario float64 `gorm:"type:decimal(10,2);not null" json:"precioUnitario"`

	// Campo calculado (no persistente en BD)
	Subtotal float64 `gorm:"-" json:"subtotal,omitempty"`
}

// Nombre correcto de la tabla en la base de datos
func (Detalle) TableName() string {
	return "detalles"
}
