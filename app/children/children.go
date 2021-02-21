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
