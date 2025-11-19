package common

type BillingValidationMessage struct {
	EventID   string      `json:"eventId"`
	EmpresaID string      `json:"empresaId"`
	NIT       string      `json:"nit"`
	Factura   interface{} `json:"factura"`
	Errores   []string    `json:"errores"`
	Resultado string      `json:"resultado"` // "VALID" o "INVALID"
}
