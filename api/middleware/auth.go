package middleware

import (
	"net/http"
	"strings"

	"github.com/PatrickChagastavares/church_backend/model"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware é a interface para geração dos middlewares
type AuthMiddleware interface {
	Public(next echo.HandlerFunc) echo.HandlerFunc
	PrivateStatic(next echo.HandlerFunc) echo.HandlerFunc
}

// newAuthMiddleware cria uma implementação da interface AuthMiddleware
func newAuthMiddleware(opts Options) AuthMiddleware {
	return &middlewareAuthImpl{
		PrivateTokenStatic: opts.PrivateTokenStatic,
	}
}

type middlewareAuthImpl struct {
	PrivateTokenStatic string
}

func (m *middlewareAuthImpl) Public(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func (m *middlewareAuthImpl) PrivateStatic(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := m.validateSession(c, true); err != nil {
			return err
		}

		return next(c)
	}
}

func (m *middlewareAuthImpl) validateSession(c echo.Context, logged bool) error {
	if logged {
		authorization := c.Request().Header.Get("Authorization")

		splitedToken := strings.Split(authorization, " ")

		if splitedToken[1] != m.PrivateTokenStatic {
			return model.NewError(http.StatusUnauthorized, "Token informado é invalido", map[string]string{
				"authorization": authorization,
			})
		}
	}

	return nil
}
