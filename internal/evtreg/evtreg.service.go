package evtreg

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	RegisterEvent(userEmail string, scheduleId int) *apperror.AppError
}

func NewService(logger *zap.Logger, repo Repository) Service {
	return &serviceImpl{
		logger,
		repo,
	}
}

var _ Service = &serviceImpl{}

type serviceImpl struct {
	logger *zap.Logger
	repo   Repository
}

func (h *serviceImpl) RegisterEvent(userEmail string, scheduleId int) *apperror.AppError {
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

	return nil
}
