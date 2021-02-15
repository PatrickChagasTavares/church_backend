package app

import (
	logger "github.com/sirupsen/logrus"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{}

	logger.Info("Initialized -> App")

	return container

}
