package doorToDoors

import (
	"net/http"
	"strconv"

	"github.com/PatrickChagastavares/church_backend/api/middleware"
	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"github.com/labstack/echo/v4"
)

// Register group DoorToDoors check
func Register(g *echo.Group, apps *app.Container, m *middleware.Middleware) {
	h := &handler{
		apps: apps,
	}

	g.POST("/create", h.create, m.Auth.PrivateStatic)
	g.GET("/list", h.listDoorToDoors, m.Auth.PrivateStatic)
	g.GET("/get/:id", h.getDoorToDoors, m.Auth.PrivateStatic)
}

type handler struct {
	apps *app.Container
}

// create swagger document
// @Description Essa rota é privada com o token valido (Bearer)
// @Description Os campos Date, NameTribe, Address e Age são obrigatorios
// @Tags DoorToDoors
// @Accept  json
// @Produce  json
// @Param door body model.DoorToDoors true "DoorToDoors"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response{data=model.DoorToDoors}
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error{message=string}}
// @Router /v1/doortodoors/create [post]
func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.DoorToDoors

	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.DoorToDoors.create.bind", err.Error())
		return model.NewError(http.StatusBadGateway, "Falha ao recuperar dados da requisição", nil)
	}
	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.DoorToDoors.create.validate", err.Error())
		return model.NewError(http.StatusBadGateway, "Requisição Inválida", nil)
	}
	child, err := h.apps.DoorToDoors.AddDoorToDoors(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: child,
	})
}

// getChildren swagger document
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags DoorToDoors
// @Accept  json
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response{data=[]model.DoorToDoors}
// @Failure 400 {object} model.Response{error=model.Error{}}
// @Failure 500 {object} model.Response{error=model.Error{}}
// @Router /v1/doortodoors/list [get]
func (h *handler) listDoorToDoors(c echo.Context) error {
	ctx := c.Request().Context()

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		logger.ErrorContext(ctx, "api.DoorToDoors.listDoorToDoors.validate", err.Error())
		return model.NewError(http.StatusBadGateway, "Page não informado", nil)
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit == 0 || limit >= 20 {
		logger.ErrorContext(ctx, "api.DoorToDoors.listDoorToDoors.validate", err.Error())
		return model.NewError(http.StatusBadGateway, "Limit não pode ser igual a zero, maior que 20 ou não informado", nil)
	}
	children, err := h.apps.DoorToDoors.ListDoorToDoors(ctx, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: children,
	})
}

// getChildren swagger document
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags DoorToDoors
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response{data=model.DoorToDoors}
// @Failure 400 {object} model.Response{error=model.Error{}}
// @Failure 500 {object} model.Response{error=model.Error{}}
// @Router /v1/doortodoors/get/{id} [get]
func (h *handler) getDoorToDoors(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.ErrorContext(ctx, "api.DoorToDoors.getDoorToDoors.validate", err.Error())
		return model.NewError(http.StatusBadGateway, "Id é zero ou não informado", nil)
	}

	child, err := h.apps.DoorToDoors.GetDoorToDoors(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: child,
	})
}
