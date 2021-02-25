package model

// Health modelo padrão para o health check
type Health struct {
	DatabaseStatus string `json:"database_status" db:"database_status"`
}
