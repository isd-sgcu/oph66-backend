package event

import (
	"context"
	"errors"
	"time"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (bool, string, *apperror.AppError)
	Set(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError
}

type cacheImpl struct {
	redis *redis.Client
}

func NewCache(redis *redis.Client) Cache {
	return &cacheImpl{
		redis,
	}
}

func (s *cacheImpl) Get(ctx context.Context, key string) (bool, string, *apperror.AppError) {
	result, err := s.redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, "", nil
	} else if err != nil {
		return false, "", apperror.InternalError
	} else {
		return true, result, nil
	}
}

func (s *cacheImpl) Set(ctx context.Context, key string, value string, expiration time.Duration) *apperror.AppError {
	err := s.redis.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return apperror.InternalError
	} else {
		return nil
	}
}
