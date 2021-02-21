package store

import (
	"github.com/PatrickChagastavares/church_backend/store/children"
	"github.com/PatrickChagastavares/church_backend/store/health"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Health   health.Store
	Children children.Store
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	Writer *gorm.DB
	Reader *gorm.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	container := &Container{
		Health:   health.NewStore(opts.Reader),
		Children: children.NewStore(opts.Reader, opts.Writer),
	}

	logger.Info("Registered -> Store")

	return container
}
