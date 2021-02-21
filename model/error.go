package model

import (
	"fmt"
	"net/http"
)

// Error tipo de erro customizado
type Error struct {
	HTTPCode int         `json:"-"`
	Message  string      `json:"message" swaggerignore:"true`
	Detail   interface{} `json:"detail,omitempty" swaggerignore:"true`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v - message: %v - detail: %v", e.HTTPCode, e.Message, e.Detail)
}

// NewError cria um novo erro
func NewError(httpCode int, message string, detail interface{}) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}

// GetHTTPCode retorna o código http do erro
func GetHTTPCode(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return http.StatusInternalServerError
	}
	return e.HTTPCode
}
