package staff

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	AttendeeStaffCheckin(userId int, departmentCode string, facultyCode string) *apperror.AppError
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		logger,
	}
}

type serviceImpl struct {
	repo   Repository
	logger *zap.Logger
}

func (s *serviceImpl) AttendeeStaffCheckin(userId int, departmentCode string, facultyCode string) *apperror.AppError {
	var checkin model.AttendeeCheckin
	checkin.UserId = userId
	checkin.DepartmentCode = departmentCode
	checkin.FacultyCode = facultyCode

	if err := s.repo.CreateCheckin(&checkin); errors.Is(err, gorm.ErrDuplicatedKey) {
		return apperror.AlreadyCheckin
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return apperror.NotFound
	} else if err != nil {
		s.logger.Error("unable to checkin", zap.Int("userId", userId), zap.String("departmentCode", departmentCode), zap.String("facultyCode", facultyCode), zap.Error(err))
		return apperror.InternalError
	}

	return nil
}
