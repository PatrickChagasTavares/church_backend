package store

import (
	"github.com/PatrickChagastavares/church_backend/store/health"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"github.com/jmoiron/sqlx"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Health health.Store
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	container := &Container{
		Health: health.NewStore(opts.Reader),
	}

	logger.Info("Registered -> Store")

	return container
}
