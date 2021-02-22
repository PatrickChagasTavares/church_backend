package children

import (
	"context"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/sirupsen/logrus"
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

	result := <-s.stores.Children.AddChild(ctx, child)
	if result.Error != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.AddChildren", result.Error.Error())

		return nil, result.Error
	}

	data, err := model.ToChild(result.Data)
	if err != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.tohealth", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *appImpl) GetChild(ctx context.Context, ID int) (*model.Child, error) {

	result := <-s.stores.Children.GetChild(ctx, ID)
	if result.Error != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.getchild", result.Error.Error())

		return nil, result.Error
	}

	data, err := model.ToChild(result.Data)
	if err != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.tochild", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *appImpl) ListChild(ctx context.Context, page, limit int) ([]*model.Child, error) {

	result := <-s.stores.Children.ListChild(ctx, limit, page)
	if result.Error != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.addchildren", result.Error.Error())

		return nil, result.Error
	}

	data, err := model.ToChildren(result.Data)
	if err != nil {
		logrus.WithContext(ctx).Error(ctx, "app.children.tochildren", err.Error())
		return nil, err
	}

	return data, nil
}
