package featureflag

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetFlag(ctx context.Context, key string) (*dto.FeatureFlagResponse, *apperror.AppError)
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		logger,
	}
}

var _ Service = &serviceImpl{}

type serviceImpl struct {
	repo   Repository
	logger *zap.Logger
}

func (h *serviceImpl) GetFlag(ctx context.Context, key string) (*dto.FeatureFlagResponse, *apperror.AppError) {
	var res model.FeatureFlag
	if err := h.repo.FindOneByKey(&res, key); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.InvalidFeatureFlagKey
	} else if err != nil {
		h.logger.Error("unable to query feature flag value from database", zap.String("key", key))
		return nil, apperror.InternalError
	} else {
		response := dto.FeatureFlagResponse{
			Key:           res.Key,
			Enabled:       res.Enabled,
			ExtraInfo:     json.RawMessage(res.ExtraInfo),
			CacheDuration: res.CacheDuration,
		}
		return &response, nil
	}
}
