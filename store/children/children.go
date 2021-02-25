package children

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Store interface para implementação do children
type Store interface {
	AddChild(ctx context.Context, child model.Child) (*model.Child, error)
	GetChild(ctx context.Context, ID int) (*model.Child, error)
	ListChild(ctx context.Context, limit, offset int) ([]*model.Child, error)
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
func (r *storeImpl) AddChild(ctx context.Context, child model.Child) (*model.Child, error) {

	result := r.writer.Create(&child)
	if result.Error != nil {
		logger.ErrorContext(ctx, "store.children.addchild", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), child)
	}
	return &child, nil

}

// GetChild busca item at table children with base in id
func (r *storeImpl) GetChild(ctx context.Context, ID int) (*model.Child, error) {

	var child model.Child
	result := r.writer.First(&child, ID)

	if result.Error == gorm.ErrRecordNotFound {
		logger.ErrorContext(ctx, "store.children.getchild.ErrRecordNotFound", result.Error.Error())
		return nil, model.NewError(http.StatusBadRequest, "Não foi encontrado informações com base no ID informado.", map[string]int{
			"Id": ID,
		})
	}

	if result.Error != nil {
		logger.ErrorContext(ctx, "store.children.getchild", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"Id": ID,
		})
	}
	return &child, nil

}

// ListChild get list of children
func (r *storeImpl) ListChild(ctx context.Context, limit, offset int) ([]*model.Child, error) {
	children := make([]*model.Child, 0)
	result := r.writer.Limit(limit).Offset(offset * limit).Find(&children)
	if result.Error != nil {
		logger.ErrorContext(ctx, "store.children.listchild", result.Error.Error())

		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"page":  offset,
			"limit": limit,
		})
	}
	return children, nil
}
