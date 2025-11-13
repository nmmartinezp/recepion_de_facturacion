package facturacion

import (
	cn "app/src/databases"
	"app/src/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

// ===========================
// VALIDACIÓN Y REGISTRO DE FACTURAS
// ===========================

// CreateFactura recibe una factura con su detalle, valida campos esenciales y la registra en la base de datos
func (s *Service) CreateFactura(factura models.Factura) (models.Factura, error) {
	// 🔍 Validaciones básicas
	if factura.CUF == "" {
		return models.Factura{}, errors.New("el campo CUF es obligatorio")
	}
	if factura.NitEmisor == "" {
		return models.Factura{}, errors.New("el NIT del emisor es obligatorio")
	}
	if factura.NitReceptor == "" {
		return models.Factura{}, errors.New("el NIT del receptor es obligatorio")
	}
	if len(factura.Detalle) == 0 {
		return models.Factura{}, errors.New("la factura debe tener al menos un detalle")
	}

	// Si la fecha de emisión no viene en el JSON, la establecemos al momento actual
	if factura.FechaEmision.IsZero() {
		factura.FechaEmision = time.Now()
	}

	// 💰 Cálculo total y validación de detalle
	var total float64 = 0
	for i, d := range factura.Detalle {
		if d.Cantidad <= 0 {
			return models.Factura{}, fmt.Errorf("el detalle #%d tiene cantidad inválida", i+1)
		}
		if d.PrecioUnitario <= 0 {
			return models.Factura{}, fmt.Errorf("el detalle #%d tiene precio inválido", i+1)
		}
		total += float64(d.Cantidad) * d.PrecioUnitario
	}

	// Validar coherencia del monto total declarado
	if factura.MontoTotal <= 0 {
		factura.MontoTotal = total
	} else if factura.MontoTotal != total {
		return models.Factura{}, fmt.Errorf("el monto total declarado (%.2f) no coincide con el calculado (%.2f)", factura.MontoTotal, total)
	}

	// 🧾 Guardar factura con sus detalles (transacción)
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&factura).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return models.Factura{}, fmt.Errorf("error al registrar la factura: %v", err)
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

// ===========================
// CONSTRUCTOR DEL SERVICIO
// ===========================

func GetService() *Service {
	db := cn.DBPOSTGRES()
	return &Service{db: db}
}
