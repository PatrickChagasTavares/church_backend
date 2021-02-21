package main

import (
	"github.com/PatrickChagastavares/church_backend/api"
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/api/swagger"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/config"
	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	logger "github.com/sirupsen/logrus"
)

func main() {

	config.Watch(func(c config.Config) {
		e := echo.New()
		e.Debug = c.GetString("env") != "prod"
		e.HideBanner = true

		e.Use(emiddleware.Logger())
		e.Use(emiddleware.BodyLimit("2M"))
		e.Use(emiddleware.Recover())
		e.Use(emiddleware.RequestID())

		dbReader, errReader := gorm.Open(mysql.Open(c.GetString("database.reader.url")), &gorm.Config{})

		if errReader != nil {
			logger.Fatal("Error ao se conectar com o database de leitura", errReader)
		}
		dbWriter, errWriter := gorm.Open(mysql.Open(c.GetString("database.writer.url")), &gorm.Config{})
		if errWriter != nil {
			logger.Fatal("Error ao se conectar com o database de leitura", errReader)
		}

		// criação dos stores com a injeção do banco de escrita e leitura
		stores := store.New(store.Options{
			Writer: dbWriter,
			Reader: dbReader,
		})

		// criação dos serviços
		apps := app.New(app.Options{
			Stores: stores,
		})

		// registros dos handlers
		api.Register(api.Options{
			Group: e.Group(""),
			Apps:  apps,

			// criação e injeção dos middlewares
			Middleware: middleware.New(middleware.Options{
				Apps:               apps,
				PrivateTokenStatic: c.GetString("PrivateTokenStatic"),
			}),
		})

		port := c.GetString("server.port")

		swagger.Register(swagger.Options{
			Group:       e.Group("/swagger"),
			AccessKey:   c.GetString("docs.key"),
			Title:       c.GetString("docs.title"),
			Description: c.GetString("docs.description"),
			Version:     c.GetString("docs.version"),
			Host:        c.GetString("docs.host"),
			BasePath:    c.GetString("docs.base-path"),
		})

		// funcão padrão pra tratamento de erros da camada http
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}

			if err := c.JSON(model.GetHTTPCode(err), model.Response{Err: err}); err != nil {
				logger.WithContext(c.Request().Context()).Info(err)
			}
		}

		go e.Start(port)

		logger.Info("Church_Backend started!")
	})
}
