package doorToDoors

import (
	"context"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/store"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
)

// App interface de health para implementação
type App interface {
	AddDoorToDoors(ctx context.Context, door model.DoorToDoors) (*model.DoorToDoors, error)
	GetDoorToDoors(ctx context.Context, ID int) (*model.DoorToDoors, error)
	ListDoorToDoors(ctx context.Context, page, limit int) ([]*model.DoorToDoors, error)
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

func (s *appImpl) AddDoorToDoors(ctx context.Context, door model.DoorToDoors) (*model.DoorToDoors, error) {

	resp, respErr := s.stores.DoorToDoors.AddDoorToDoors(ctx, door)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.DoorToDoors.AddDoorToDoors", respErr.Error())
		return nil, respErr
	}

	return resp, nil

}

func (s *appImpl) GetDoorToDoors(ctx context.Context, ID int) (*model.DoorToDoors, error) {

	resp, respErr := s.stores.DoorToDoors.GetDoorToDoors(ctx, ID)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.DoorToDoors.GetDoorToDoors", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}

func (s *appImpl) ListDoorToDoors(ctx context.Context, page, limit int) ([]*model.DoorToDoors, error) {

	resp, respErr := s.stores.DoorToDoors.ListDoorToDoors(ctx, limit, page)
	if respErr != nil {
		logger.ErrorContext(ctx, "app.DoorToDoors.ListDoorToDoors", respErr.Error())
		return nil, respErr
	}

	return resp, nil
}
