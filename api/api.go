package api

import (
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	v1 "github.com/PatrickChagastavares/church_backend/api/v1"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Group      *echo.Group
	Apps       *app.Container
	Middleware *middleware.Middleware
}

// Register register routes
func Register(opts Options) {
	v1.Register(opts.Group, opts.Apps, opts.Middleware)

	logger.Info("Initialized -> Api")

}
