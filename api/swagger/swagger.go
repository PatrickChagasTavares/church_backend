package swagger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	docs "github.com/PatrickChagastavares/church_backend/docs"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
)

// Options struct de opções para a criação de uma instancia do swagger
type Options struct {
	Group       *echo.Group
	AccessKey   string
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
}

// Register group item check
func Register(opts Options) {

	docs.SwaggerInfo.Title = opts.Title
	docs.SwaggerInfo.Description = opts.Description
	docs.SwaggerInfo.Version = opts.Version
	docs.SwaggerInfo.Host = opts.Host
	docs.SwaggerInfo.BasePath = opts.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	opts.Group.GET("/:key", func(c echo.Context) error {
		key := c.Param("key")
		if key != opts.AccessKey {
			return nil
		}
		return c.Redirect(http.StatusFound, "/swagger/"+key+"/index.html")
	})

	opts.Group.GET("/:key/*", func(c echo.Context) error {
		key := c.Param("key")

		if key != opts.AccessKey {
			return c.JSON(
				http.StatusUnauthorized,
				notAuthorized{http.StatusUnauthorized, "You are not wellcome here"},
			)
		}

		return echoSwagger.WrapHandler(c)
	})

	logger.Info("Swagger is initializing...")
}
