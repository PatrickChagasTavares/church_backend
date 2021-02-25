package app

import (
	"github.com/PatrickChagastavares/church_backend/app/children"
	"github.com/PatrickChagastavares/church_backend/app/doorToDoors"
	"github.com/PatrickChagastavares/church_backend/app/health"
	"github.com/PatrickChagastavares/church_backend/app/social"
	"github.com/PatrickChagastavares/church_backend/store"
	logger "github.com/sirupsen/logrus"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health      health.App
	Children    children.App
	DoorToDoors doorToDoors.App
	Social      social.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		Health:      health.NewApp(opts.Stores),
		Children:    children.NewApp(opts.Stores),
		DoorToDoors: doorToDoors.NewApp(opts.Stores),
		Social:      social.NewApp(opts.Stores),
	}

	logger.Info("Initialized -> App")

	return container

}
