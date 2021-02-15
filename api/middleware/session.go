package middleware

import (
	"net/http"
	"strings"

	"github.com/PatrickChagastavares/church_backend/app"
	"github.com/PatrickChagastavares/church_backend/model"

	"github.com/labstack/echo/v4"
)

// SessionMiddleware é a interface para geração dos middlewares
type SessionMiddleware interface {
	InjectSession(next echo.HandlerFunc) echo.HandlerFunc
}

// newSessionMiddleware cria uma implementação da interface SessionMiddleware
func newSessionMiddleware(opts Options) SessionMiddleware {
	return &middlewareSessionImpl{
		apps:               opts.Apps,
		PrivateTokenStatic: opts.PrivateTokenStatic,
	}
}

type middlewareSessionImpl struct {
	apps               *app.Container
	PrivateTokenStatic string
}

func (m *middlewareSessionImpl) InjectSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization != "" {
			splitedToken := strings.Split(authorization, " ")
			if len(splitedToken) != 2 {
				return model.NewError(http.StatusUnauthorized, "não foi possível decodificar o token", map[string]string{
					"authorization": authorization,
				})
			}

			if splitedToken[1] != m.PrivateTokenStatic {
				return model.NewError(http.StatusUnauthorized, "Token informado é invalido", map[string]string{
					"authorization": authorization,
				})
			}
		}

		return next(c)
	}
}
