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
	GetChild(ctx context.Context, ID int) do.ChanResult
	ListChild(ctx context.Context, limit, offset int) do.ChanResult
}

// NewStore cria uma nova instancia do repositorio de children
func NewStore(reader, writer *gorm.DB) Store {
	return &storeImpl{reader, writer}
}

type storeImpl struct {
	reader *gorm.DB
	writer *gorm.DB
}

// AddChild create item at table children
func (r *storeImpl) AddChild(ctx context.Context, child model.Child) do.ChanResult {
	return do.Do(func(res *do.Result) {
		result := r.writer.Create(&child)
		if result.Error != nil {
			logger.ErrorContext(ctx, "store.children.addchild", result.Error.Error())
			res.Error = model.NewError(http.StatusInternalServerError, result.Error.Error(), child)
			return
		}
		res.Data = &child
	})
}

// GetChild busca item at table children with base in id
func (r *storeImpl) GetChild(ctx context.Context, ID int) do.ChanResult {
	return do.Do(func(res *do.Result) {
		var child model.Child
		result := r.writer.First(&child, ID)
		logger.Info(result.RowsAffected)

		if result.RowsAffected == 0 {
			logger.ErrorContext(ctx, "store.children.getchild.rowsaffected", result.Error.Error())
			res.Error = model.NewError(http.StatusBadRequest, result.Error.Error(), map[string]int{
				"Id": ID,
			})
			return
		}
		if result.Error != nil {
			logger.ErrorContext(ctx, "store.children.getchild", result.Error.Error())
			res.Error = model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
				"Id": ID,
			})
			return
		}
		res.Data = &child
	})
}

// ListChild get list of children
func (r *storeImpl) ListChild(ctx context.Context, limit, offset int) do.ChanResult {
	return do.Do(func(res *do.Result) {
		children := make([]*model.Child, 0)
		result := r.writer.Limit(limit).Offset(offset * limit).Find(&children)
		if result.Error != nil {
			logger.ErrorContext(ctx, "store.children.listchild", result.Error.Error())
			res.Error = model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
				"page":  offset,
				"limit": limit,
			})
			return
		}
		res.Data = children
	})
}
