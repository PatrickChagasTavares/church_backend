package v1

import (
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/api/v1/health"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group, apps *app.Container, middleware *middleware.Middleware) {
	v1 := g.Group("/v1", middleware.Session.InjectSession)

	health.Register(v1.Group("/health"), apps, middleware)
	// cpf.Register(v1.Group("/cpf"), apps, middleware)
}
