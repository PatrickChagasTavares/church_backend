package children

import (
	"net/http"

	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"github.com/labstack/echo/v4"
)

// Register group health check
func Register(g *echo.Group, apps *app.Container, m *middleware.Middleware) {
	h := &handler{
		apps: apps,
	}

	g.POST("/create", h.create, m.Auth.PrivateStatic)
	// g.GET("/check", h.check, m.Auth.PrivateStatic)
}

type handler struct {
	apps *app.Container
}

// create swagger document
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags children
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response{data=model.Child}
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error{message=string}}
// @Router /v1/children/create [post]
func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.Child

	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.children.create.bind", err.Error())
		return model.NewError(http.StatusBadGateway, "Falha ao recuperar dados da requisição", nil)
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.children.create.validate", err.Error())
		return model.NewError(http.StatusBadGateway, "Requisição Inválida", nil)
	}
	children, err := h.apps.Children.AddChildren(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: children,
	})
}
