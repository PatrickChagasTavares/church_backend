package health

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Store interface para implementação do health
type Store interface {
	Ping(ctx context.Context) (*model.Health, error)
	Check(ctx context.Context) (*model.Health, error)
}

// NewStore cria uma nova instancia do repositorio de health
func NewStore(reader *gorm.DB) Store {
	return &storeImpl{reader}
}

type storeImpl struct {
	reader *gorm.DB
}

// Ping checa se o banco está online
func (r *storeImpl) Ping(ctx context.Context) (*model.Health, error) {
	_, err := r.reader.DB()
	if err != nil {
		logger.ErrorContext(ctx, "store.health.ping", err)

		return nil, model.NewError(http.StatusInternalServerError, err.Error(), nil)
	}
	return &model.Health{DatabaseStatus: "OK"}, nil
}

// Check checa se o banco está com status OK
func (r *storeImpl) Check(ctx context.Context) (*model.Health, error) {
	data := new(model.Health)

	query := r.reader.WithContext(ctx).Raw(`SELECT 'DB OK' AS database_status`).Row()
	err := query.Scan(data)
	if err != nil {
		logger.ErrorContext(ctx, "store.health.check", err.Error())
		return nil, model.NewError(http.StatusInternalServerError, err.Error(), nil)
	}

	return data, nil
}
