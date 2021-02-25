package doorToDoors

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Store interface para implementação do DoorToDoors
type Store interface {
	AddDoorToDoors(ctx context.Context, door model.DoorToDoors) (*model.DoorToDoors, error)
	GetDoorToDoors(ctx context.Context, ID int) (*model.DoorToDoors, error)
	ListDoorToDoors(ctx context.Context, limit, offset int) ([]*model.DoorToDoors, error)
}

// NewStore cria uma nova instancia do repositorio de DoorToDoors
func NewStore(reader, writer *gorm.DB) Store {
	return &storeImpl{reader, writer}
}

type storeImpl struct {
	reader *gorm.DB
	writer *gorm.DB
}

func (s *storeImpl) AddDoorToDoors(ctx context.Context, door model.DoorToDoors) (*model.DoorToDoors, error) {

	result := s.writer.Create(&door)
	if result.Error != nil {
		logger.ErrorContext(ctx, "store.children.addchild", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), door)
	}
	return &door, nil
}

func (s *storeImpl) GetDoorToDoors(ctx context.Context, ID int) (*model.DoorToDoors, error) {

	var door model.DoorToDoors
	result := s.reader.First(&door, ID)

	if result.Error == gorm.ErrRecordNotFound {
		logger.ErrorContext(ctx, "store.doortodoors.getchild.ErrRecordNotFound", result.Error.Error())
		return nil, model.NewError(http.StatusBadRequest, "Não foi encontrado informações com base no ID informado.", map[string]int{
			"Id": ID,
		})
	}

	if result.Error != nil {
		logger.ErrorContext(ctx, "store.doortodoors.getchild", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"Id": ID,
		})
	}
	return &door, nil
}

func (s *storeImpl) ListDoorToDoors(ctx context.Context, limit, offset int) ([]*model.DoorToDoors, error) {
	doors := make([]*model.DoorToDoors, 0)
	result := s.writer.Limit(limit).Offset(offset * limit).Find(&doors)

	if result.Error != nil {
		logger.ErrorContext(ctx, "store.doortodoors.listchild", result.Error.Error())

		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"page":  offset,
			"limit": limit,
		})
	}
	return doors, nil
}
