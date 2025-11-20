package models

type Detalle struct {
	ID             uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FacturaID      uint64  `gorm:"not null;index" json:"factura_id"`
	Descripcion    string  `gorm:"type:varchar(200);not null" json:"descripcion"`
	Cantidad       int     `gorm:"not null" json:"cantidad"`
	PrecioUnitario float64 `gorm:"type:numeric(10,2);not null" json:"precio_unitario"`

	// Relación inversa
	Factura *Factura `gorm:"foreignKey:FacturaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"factura"`
}

// Nombre correcto de la tabla en la base de datos
func (Detalle) TableName() string {
	return "detalles"
}
