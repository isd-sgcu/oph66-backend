package featureflag

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache interface {
	// Return nil, nil when no error and also not found
	Get(ctx context.Context, key string) (json.RawMessage, *apperror.AppError)
	Set(ctx context.Context, key string, value *dto.FeatureFlagResponse, cacheDuration time.Duration) *apperror.AppError
}

func NewCache(client *redis.Client, logger *zap.Logger) Cache {
	return &cacheImpl{
		client,
		logger,
	}
}

const cacheKeyPrefix = "featureflag-"

var _ Cache = &cacheImpl{}

type cacheImpl struct {
	client *redis.Client
	logger *zap.Logger
}

func (c *cacheImpl) Get(ctx context.Context, key string) (json.RawMessage, *apperror.AppError) {
	prefixedKey := cacheKeyPrefix + key
	result, err := c.client.Get(ctx, prefixedKey).Result()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		c.logger.Error("unable to get feature flag cache", zap.String("key", key), zap.Error(err))
		return nil, apperror.InternalError
	} else {
		return json.RawMessage(result), nil
	}
}

func (c *cacheImpl) Set(ctx context.Context, key string, value *dto.FeatureFlagResponse, cacheDuration time.Duration) *apperror.AppError {
	prefixedKey := cacheKeyPrefix + key
	raw, err := json.Marshal(value)

	if err != nil {
		c.logger.Error("unable to marshal feature flag", zap.String("key", key), zap.Any("value", value))
		return apperror.InternalError
	}

	rawString := string(raw)

	if err = c.client.Set(ctx, prefixedKey, rawString, cacheDuration).Err(); err != nil {
		c.logger.Error("unable to set feature flag cache", zap.String("key", key), zap.Error(err), zap.ByteString("raw", raw))
		return apperror.InternalError
	}
	return nil
}
