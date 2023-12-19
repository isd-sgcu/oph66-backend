package event

import (
	"context"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Service interface {
	GetAllEvents(ctx context.Context) ([]Event, *apperror.AppError)
}

func NewService(repo Repository, redis *redis.Client, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		redis,
		logger,
	}
}

type serviceImpl struct {
	repo   Repository
	redis  *redis.Client
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
