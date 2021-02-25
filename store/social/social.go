package social

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/logger"
	"gorm.io/gorm"
)

// Store interface para implementação do DoorToDoors
type Store interface {
	AddSocial(ctx context.Context, social model.Social) (*model.Social, error)
	GetSocial(ctx context.Context, ID int) (*model.Social, error)
	ListSocial(ctx context.Context, limit, offset int) ([]*model.Social, error)
}

// NewStore cria uma nova instancia do repositorio de DoorToDoors
func NewStore(reader, writer *gorm.DB) Store {
	return &storeImpl{reader, writer}
}

type storeImpl struct {
	reader *gorm.DB
	writer *gorm.DB
}

func (s *storeImpl) AddSocial(ctx context.Context, social model.Social) (*model.Social, error) {
	result := s.writer.Create(&social)
	if result.Error != nil {
		logger.ErrorContext(ctx, "store.social.addsocial", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), social)
	}
	return &social, nil
}

func (s *storeImpl) GetSocial(ctx context.Context, ID int) (*model.Social, error) {
	var social model.Social
	result := s.reader.First(&social, ID)

	if result.Error == gorm.ErrRecordNotFound {
		logger.ErrorContext(ctx, "store.social.getsocial.ErrRecordNotFound", result.Error.Error())
		return nil, model.NewError(http.StatusBadRequest, "Não foi encontrado informações com base no ID informado.", map[string]int{
			"Id": ID,
		})
	}

	if result.Error != nil {
		logger.ErrorContext(ctx, "store.social.getchild", result.Error.Error())
		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"Id": ID,
		})
	}
	return &social, nil
}

func (s *storeImpl) ListSocial(ctx context.Context, limit, offset int) ([]*model.Social, error) {
	social := make([]*model.Social, 0)
	result := s.writer.Limit(limit).Offset(offset * limit).Find(&social)

	if result.Error != nil {
		logger.ErrorContext(ctx, "store.social.listsocial", result.Error.Error())

		return nil, model.NewError(http.StatusInternalServerError, result.Error.Error(), map[string]int{
			"page":  offset,
			"limit": limit,
		})
	}
	return social, nil
}
