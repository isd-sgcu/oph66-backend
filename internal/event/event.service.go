package event

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetAllEvents() ([]dto.Event, *apperror.AppError)
	GetEventById(eventId string) (dto.Event, *apperror.AppError)
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		logger,
	}
}

type serviceImpl struct {
	repo   Repository
	logger *zap.Logger
}

func (s *serviceImpl) GetAllEvents() ([]dto.Event, *apperror.AppError) {
	var query []model.Event
	err := s.repo.GetAllEvents(&query)
	if err != nil {
		s.logger.Error("could not retrieve events from database", zap.Error(err))
		return []dto.Event{}, apperror.InternalError
	}

	results := make([]dto.Event, 0, len(query))
	for _, r := range query {
		results = append(results, EventModelToDTO(&r))
	}

	return results, nil
}

func (s *serviceImpl) GetEventById(eventId string) (dto.Event, *apperror.AppError) {
	var mEvent model.Event
	err := s.repo.GetEventById(&mEvent, eventId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("could not retrieve event with the specified event id", zap.String("eventId", eventId), zap.Error(err))
		return dto.Event{}, apperror.InvalidEventId
	} else if err != nil {
		s.logger.Error("could not retrieve event from database", zap.String("eventId", eventId), zap.Error(err))
		return dto.Event{}, apperror.InternalError
	}

	result := EventModelToDTO(&mEvent)

	return result, nil
}
