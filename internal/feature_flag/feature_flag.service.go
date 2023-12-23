package featureflag

import (
	"context"
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetFlag(ctx context.Context, key string) (*FeatureFlag, *apperror.AppError)
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

func (h *serviceImpl) GetFlag(ctx context.Context, key string) (*FeatureFlag, *apperror.AppError) {
	var res FeatureFlag
	if err := h.repo.FindOneByKey(&res, key); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.InvalidFeatureFlagKey
	} else if err != nil {
		h.logger.Error("unable to query feature flag value from database", zap.String("key", key))
		return nil, apperror.InternalError
	} else {
		return &res, nil
	}
}
