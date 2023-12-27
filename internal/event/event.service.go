package event

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetAllEvents() ([]model.Event, *apperror.AppError)
	GetEventById(eventId string) (model.Event, *apperror.AppError)
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

func (s *serviceImpl) GetAllEvents() ([]model.Event, *apperror.AppError) {
	results := []model.Event{}
	err := s.repo.GetAllEvents(&results)
	if err != nil {
		s.logger.Error("could not retrieve events from database", zap.Error(err))
		return []model.Event{}, apperror.InternalError
	}

	return results, nil
}

func (s *serviceImpl) GetEventById(eventId string) (model.Event, *apperror.AppError) {
	result := model.Event{}
	err := s.repo.GetEventById(&result, eventId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("could not retrieve event with the specified event id", zap.String("eventId", eventId), zap.Error(err))
		return model.Event{}, apperror.InvalidEventId
	} else if err != nil {
		s.logger.Error("could not retrieve event from database", zap.String("eventId", eventId), zap.Error(err))
		return model.Event{}, apperror.InternalError
	}

	return result, nil
}
