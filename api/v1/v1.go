package v1

import (
	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/api/v1/children"
	"github.com/PatrickChagastavares/church_backend/api/v1/doorToDoors"
	"github.com/PatrickChagastavares/church_backend/api/v1/health"
	"github.com/PatrickChagastavares/church_backend/api/v1/social"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/labstack/echo/v4"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Middleware) {
	v1 := g.Group("/v1", middleware.Session.InjectSession)

	health.Register(v1.Group("/health"), apps, middleware)
	children.Register(v1.Group("/children"), apps, middleware)
	doorToDoors.Register(v1.Group("/doortodoors"), apps, middleware)
	social.Register(v1.Group("/social"), apps, middleware)
}
