package model

// Child modelo padrão para o Child check
type Child struct {
	ID        int32  `json:"Id,omitempty" db:"id" gorm:"primaryKey;autoIncrement:true" swaggerignore:"true"`
	Date      int64  `json:"Date" db:"date"  gorm:"not null" validate:"required"`
	Total     int32  `json:"Total" db:"total" gorm:"not null" validate:"required"`
	Notes     string `json:"Notes" db:"notes"`
	CreatedAt int64  `json:"CreatedAt,omitempty" db:"created_at" swaggerignore:"true"`
	UpdatedAt int64  `json:"UpdatedAt,omitempty" db:"updated_at" swaggerignore:"true"`
	DeletedAt int64  `json:"DeletedAt,omitempty" db:"deleted_at" swaggerignore:"true"`
}
