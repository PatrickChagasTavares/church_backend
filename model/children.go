package model

import (
	"net/http"
)

// Child modelo padrão para o Child check
type Child struct {
	ID        int32  `json:"Id,omitempty" db:"id" gorm:"primaryKey;autoIncrement:true"`
	Date      int64  `json:"Date" db:"date"  gorm:"not null" validate:"required"`
	Total     int32  `json:"Total" db:"total" gorm:"not null" validate:"required"`
	Notes     string `json:"Notes" db:"notes"`
	CreatedAt int64  `json:"CreatedAt,omitempty" db:"created_at"`
	UpdatedAt int64  `json:"UpdatedAt,omitempty" db:"updated_at"`
	DeletedAt int64  `json:"DeletedAt,omitempty" db:"deleted_at"`
}

// ToChild converte uma interface{} para *Chidren
func ToChild(data interface{}) (*Child, error) {
	value, ok := data.(*Child)
	if !ok {
		return nil, NewError(http.StatusInternalServerError, "não foi possível converter interface{} para *Child", nil)
	}
	return value, nil
}

// ToChildren converte uma interface{} para []*Children
func ToChildren(data interface{}) ([]*Child, error) {
	value, ok := data.([]*Child)
	if !ok {
		return nil, NewError(http.StatusInternalServerError, "não foi possível converter interface{} para []*Child", nil)
	}
	return value, nil
}
