package model

// Health modelo padrão para o health check
type Health struct {
	Version        string `json:"version"`
	ServerStatedAt string `json:"server_started_at"`
	DatabaseStatus string `json:"database_status" db:"database_status"`
}
