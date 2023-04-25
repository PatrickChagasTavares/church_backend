package main

import (
	"fmt"
	"os"

	"github.com/PatrickChagastavares/church_backend/api"
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/api/swagger"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/PatrickChagastavares/church_backend/utils/validator"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	logger "github.com/sirupsen/logrus"
)

// main configure swagger
// method of use bearer token in requests
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load(".env")

	// config.Watch(func(c config.Config) {
	e := echo.New()
	e.Validator = validator.New()
	e.Debug = os.Getenv("ENV") != "prod"
	e.HideBanner = true

	e.Use(emiddleware.Logger())
	e.Use(emiddleware.BodyLimit("2M"))
	e.Use(emiddleware.Recover())
	e.Use(emiddleware.RequestID())

	url := os.Getenv("database.reader.url")
	fmt.Println("url ->> ", url)

	dbReader, errReader := gorm.Open(postgres.Open(os.Getenv("database.reader.url")), &gorm.Config{})

	if errReader != nil {
		logger.Fatal("Error ao se conectar com o database de leitura", errReader)
	}

	dbWriter, errWriter := gorm.Open(postgres.Open(os.Getenv("database.writer.url")), &gorm.Config{})
	if errWriter != nil {
		logger.Fatal("Error ao se conectar com o database de escrita", errReader)
	}

	dbWriter.Migrator().AutoMigrate(model.Child{}, model.DoorToDoors{}, model.Social{})

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
			PrivateTokenStatic: os.Getenv("SESSION_SECRET"),
		}),
	})

	swagger.Register(swagger.Options{
		Group:       e.Group("/swagger"),
		AccessKey:   os.Getenv("docs.key"),
		Title:       os.Getenv("docs.title"),
		Description: os.Getenv("docs.description"),
		Version:     os.Getenv("docs.version"),
		Host:        os.Getenv("docs.host"),
		BasePath:    os.Getenv("docs.baseURL"),
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

	logger.Info("Church_Backend started!")
	e.Start(os.Getenv("PORT"))

	// })
}
