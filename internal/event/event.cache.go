package event

import (
	"context"
	"errors"
	"time"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache interface {
	Get(ctx context.Context, key string) (bool, string, *apperror.AppError)
	Set(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError
}

type cacheImpl struct {
	redis  *redis.Client
	logger *zap.Logger
}

func NewCache(redis *redis.Client, logger *zap.Logger) Cache {
	return &cacheImpl{
		redis,
		logger,
	}
}

func (s *cacheImpl) Get(ctx context.Context, key string) (bool, string, *apperror.AppError) {
	result, err := s.redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, "", nil
	} else if err != nil {
		s.logger.Error("could not retrieve data from redis", zap.String("key", key), zap.Error(err))
		return false, "", apperror.InternalError
	} else {
		return true, result, nil
	}
}

func (s *cacheImpl) Set(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError {
	err := s.redis.Set(ctx, key, value, expiration).Err()
	if err != nil {
		s.logger.Error("could not set key value pair on redis", zap.String("key", key), zap.String("value", value), zap.Error(err))
		return apperror.InternalError
	} else {
		return nil
	}
}
