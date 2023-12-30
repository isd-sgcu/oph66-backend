package evtreg

import (
	"context"
	"errors"
	"fmt"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/event"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	RegisterEvent(ctx context.Context, userEmail string, scheduleId int) *apperror.AppError
}

func NewService(logger *zap.Logger, repo Repository, cache event.Cache) Service {
	return &serviceImpl{
		logger,
		repo,
		cache,
	}
}

var _ Service = &serviceImpl{}

type serviceImpl struct {
	logger *zap.Logger
	repo   Repository
	cache  event.Cache
}

func (h *serviceImpl) RegisterEvent(ctx context.Context, userEmail string, scheduleId int) *apperror.AppError {
	var user model.User
	if err := h.repo.GetUserWithEventRegistrationByEmail(&user, userEmail); errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.UserNotFound
	} else if err != nil {
		h.logger.Error("unable to query user by email", zap.String("email", userEmail), zap.Int("scheduleId", scheduleId))
		return apperror.InternalError
	}

	var schedule model.Schedule
	if err := h.repo.GetScheduleById(&schedule, scheduleId); errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.ScheduleNotFound
	} else if err != nil {
		h.logger.Error("unable to get schedule by id", zap.String("email", userEmail), zap.Int("scheduleId", scheduleId))
		return apperror.InternalError
	}

	for _, regEvt := range user.RegisteredEvents {
		if regEvt.Schedule.Period == schedule.Period {
			return apperror.DuplicatePeriod
		}
	}

	if err := h.repo.RegisterEvent(user.Id, scheduleId); errors.Is(err, apperror.ScheduleFull) {
		return apperror.ScheduleFull
	} else if err != nil {
		h.logger.Error("unable to register event", zap.Int("userId", user.Id), zap.Int("scheduleId", scheduleId))
		return apperror.InternalError
	}

	keys := []string{fmt.Sprintf("get_event_by_id-%v", schedule.EventId), "get_all_events"}
	if apperr := h.cache.Del(ctx, keys...); apperr != nil {
		return apperr
	}

	return nil
}
