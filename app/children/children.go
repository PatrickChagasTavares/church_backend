package children

import (
	"context"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
)

// App interface de health para implementação
type App interface {
	AddChildren(ctx context.Context, child model.Child) (*model.Child, error)
	GetChild(ctx context.Context, ID int) (*model.Child, error)
	ListChild(ctx context.Context, page, limit int) ([]*model.Child, error)
}

// NewApp cria uma nova instancia do serviço de health
func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores *store.Container
}

func (s *appImpl) AddChildren(ctx context.Context, child model.Child) (*model.Child, error) {

	resp, respErr := s.stores.Children.AddChild(ctx, child)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.children.AddChildren", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}

func (s *appImpl) GetChild(ctx context.Context, ID int) (*model.Child, error) {

	resp, respErr := s.stores.Children.GetChild(ctx, ID)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.children.getchild", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}

func (s *appImpl) ListChild(ctx context.Context, page, limit int) ([]*model.Child, error) {

	resp, respErr := s.stores.Children.ListChild(ctx, limit, page)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.children.addchildren", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}
