package event

import (
	"context"
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetAllEvents(ctx context.Context) ([]Event, *apperror.AppError)
	GetEventById(ctx context.Context, eventId string) (Event, *apperror.AppError)
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

func (s *serviceImpl) GetAllEvents(ctx context.Context) ([]Event, *apperror.AppError) {
	results := []Event{}
	err := s.repo.GetAllEvents(&results)
	if err != nil {
		return []Event{}, apperror.InternalError
	}

	return results, nil
}

func (s *serviceImpl) GetEventById(ctx context.Context, eventId string) (Event, *apperror.AppError) {
	result := Event{}
	err := s.repo.GetEventById(&result, eventId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Event{}, apperror.InvalidEventId
	} else if err != nil {
		return Event{}, apperror.InternalError
	}

	return result, nil
}
