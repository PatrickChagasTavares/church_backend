package middleware

import "github.com/PatrickChagastavares/church_backend/app"

// Options struct de opções para a criação de uma instancia dos middlewares
type Options struct {
	Apps               *app.Container
	PrivateTokenStatic string
}

// Middleware é um container para as implementações
type Middleware struct {
	Auth    AuthMiddleware
	Session SessionMiddleware
}

// New cria uma nova instancia dos middlewares injetando os serviços
func New(opts Options) *Middleware {
	return &Middleware{
		Auth:    newAuthMiddleware(opts),
		Session: newSessionMiddleware(opts),
	}
}
