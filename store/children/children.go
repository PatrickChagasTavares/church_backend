package children

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/do"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Store interface para implementação do children
type Store interface {
	AddChild(ctx context.Context, child model.Child) do.ChanResult
}

// NewStore cria uma nova instancia do repositorio de children
func NewStore(reader, writer *gorm.DB) Store {
	return &storeImpl{reader, writer}
}

type storeImpl struct {
	reader *gorm.DB
	writer *gorm.DB
}

// Ping checa se o banco está online
func (r *storeImpl) AddChild(ctx context.Context, child model.Child) do.ChanResult {
	return do.Do(func(res *do.Result) {
		result := r.writer.Table("children").Create(child)
		if result.Error != nil {
			logger.ErrorContext(ctx, "store.health.ping", result.Error.Error())
			res.Error = model.NewError(http.StatusInternalServerError, result.Error.Error(), nil)
			return
		}
		res.Data = child
	})
}
