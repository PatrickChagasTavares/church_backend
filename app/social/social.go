package social

import (
	"context"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
)

// App interface de health para implementação
type App interface {
	AddSocial(ctx context.Context, social model.Social) (*model.Social, error)
	GetSocial(ctx context.Context, ID int) (*model.Social, error)
	ListSocial(ctx context.Context, page, limit int) ([]*model.Social, error)
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

func (s *appImpl) AddSocial(ctx context.Context, social model.Social) (*model.Social, error) {

	resp, respErr := s.stores.Social.AddSocial(ctx, social)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.social.AddSocial", respErr.Error())
		return nil, respErr
	}

	return resp, nil

}

func (s *appImpl) GetSocial(ctx context.Context, ID int) (*model.Social, error) {

	resp, respErr := s.stores.Social.GetSocial(ctx, ID)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.social.GetSocial", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}

func (s *appImpl) ListSocial(ctx context.Context, page, limit int) ([]*model.Social, error) {

	resp, respErr := s.stores.Social.ListSocial(ctx, limit, page)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.social.ListSocial", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}
