package featureflag

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
	GetFlag(ctx context.Context, key string) (bool, *apperror.AppError)
}

func NewService(db *gorm.DB, redis *redis.Client, logger *zap.Logger) Service {
	return &serviceImpl{
		db,
		redis,
		logger,
	}
}

type serviceImpl struct {
	db     *gorm.DB
	redis  *redis.Client
	logger *zap.Logger
}

func (h *serviceImpl) GetFlag(ctx context.Context, key string) (bool, *apperror.AppError) {
	res, err := h.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		var res FeatureFlag
		if err := h.db.First(&res).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return false, apperror.InvalidFeatureFlagKey
		} else if err != nil {
			h.logger.Error("unable to query feature flag value from database", zap.String("key", key))
			return false, apperror.InternalError
		}

		if err := h.setCache(ctx, key, res.Value, res.CacheDuration); err != nil {
			h.logger.Error("unable to set cache", zap.String("key", key), zap.Bool("value", res.Value), zap.Int("CacheDuration", res.CacheDuration))
			return false, apperror.InternalError
		}

		return res.Value, nil
	} else if err == nil {
		switch res {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			h.logger.Error("invalid feature flag value from redis", zap.String("key", key))
			return false, apperror.InternalError
		}
	} else {
		h.logger.Error("unable to retrieve value from redis", zap.String("key", key), zap.Error(err))
		return false, apperror.InternalError
	}
}

func (h *serviceImpl) setCache(ctx context.Context, key string, value bool, cacheDuration int) *apperror.AppError {
	var stringValue string
	if value {
		stringValue = "true"
	} else {
		stringValue = "false"
	}

	if err := h.redis.Set(ctx, key, stringValue, time.Duration(cacheDuration*int(time.Second))).Err(); err != nil {
		h.logger.Error("unable to put flag on redis", zap.String("key", key), zap.Bool("value", value), zap.Int("cacheDuration", cacheDuration), zap.Error(err))
		return apperror.InternalError
	}
	return nil
}
