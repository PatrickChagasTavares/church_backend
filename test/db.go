package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDB retorna uma instancia do db mocado
func GetDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	return gormDB, mock
}

// NewRows retorna um modelo para adicionar rows
func NewRows(columns ...string) *sqlmock.Rows {
	return sqlmock.NewRows(columns)
}
