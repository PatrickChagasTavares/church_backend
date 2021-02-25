package health

import (
	"context"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/sirupsen/logrus"
)

// App interface de health para implementação
type App interface {
	Ping(ctx context.Context) (*model.Health, error)
	Check(ctx context.Context) (*model.Health, error)
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

func (s *appImpl) Ping(ctx context.Context) (*model.Health, error) {
	resp, respErr := s.stores.Health.Ping(ctx)
	if respErr != nil {
		logrus.WithContext(ctx).Error(ctx, "app.health.ping", respErr.Error())

		return nil, respErr
	}

	return resp, nil
}

func (s *appImpl) Check(ctx context.Context) (*model.Health, error) {
	resp, respErr := s.stores.Health.Check(ctx)
	if respErr != nil {
		logrus.WithContext(ctx).Error("app.health.check", respErr.Error())

		return nil, respErr
	}

	return resp, nil
}
