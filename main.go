package main

import (
	"github.com/PatrickChagastavares/church_backend/api"
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/api/swagger"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/config"
	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"

	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {

	config.Watch(func(c config.Config) {
		e := echo.New()
		e.HideBanner = true

		e.Use(emiddleware.Logger())
		e.Use(emiddleware.BodyLimit("2M"))
		e.Use(emiddleware.Recover())
		e.Use(emiddleware.RequestID())

		// criação dos serviços
		apps := app.New(app.Options{})

		// registros dos handlers
		api.Register(api.Options{
			Group: e.Group(""),
			Apps:  apps,

			// criação e injeção dos middlewares
			Middleware: middleware.New(middleware.Options{
				Apps: apps,
			}),
		})

		port := c.GetString("server.port")

		swagger.Register(swagger.Options{
			Port:      port,
			Group:     e.Group("/swagger"),
			AccessKey: c.GetString("docs.key"),
		})

		// funcão padrão pra tratamento de erros da camada http
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}

			if err := c.JSON(model.GetHTTPCode(err), model.Response{Err: err}); err != nil {
				logger.ErrorContext(c.Request().Context(), err)
			}
		}

		go e.Start(port)

		logger.Info("Church_Backend started!")
	})
}
