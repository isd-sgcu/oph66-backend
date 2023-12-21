package event

import (
	"context"
	"errors"
	"time"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetAllEvents(ctx context.Context) ([]Event, *apperror.AppError)
	GetEventById(ctx context.Context, eventId string) (Event, *apperror.AppError)
	GetEventCache(ctx context.Context, key string) (bool, string, *apperror.AppError)
	SetEventCache(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError
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

func (s *serviceImpl) GetEventCache(ctx context.Context, key string) (bool, string, *apperror.AppError) {
	result, err := s.redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, "", nil
	} else if err != nil {
		return false, "", apperror.InternalError
	} else {
		return true, result, nil
	}
}

func (s *serviceImpl) SetEventCache(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError {
	err := s.redis.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return apperror.InternalError
	} else {
		return nil
	}
}
