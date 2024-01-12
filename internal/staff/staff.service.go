package staff

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	AttendeeFacultyStaffCheckin(userId int, departmentCode string, facultyCode string) *apperror.AppError
	AttendeeCentralStaffCheckin(userId int) *apperror.AppError
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

func (s *serviceImpl) AttendeeFacultyStaffCheckin(userId int, departmentCode string, facultyCode string) *apperror.AppError {
	var checkin model.AttendeeFacultyCheckin
	checkin.UserId = userId
	checkin.DepartmentCode = departmentCode
	checkin.FacultyCode = facultyCode

	if err := s.repo.CreateFacultyCheckin(&checkin); errors.Is(err, gorm.ErrDuplicatedKey) {
		return apperror.AlreadyCheckin
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return apperror.NotFound
	} else if err != nil {
		s.logger.Error("unable to faculty checkin", zap.Int("userId", userId), zap.String("departmentCode", departmentCode), zap.String("facultyCode", facultyCode), zap.Error(err))
		return apperror.InternalError
	}

	return nil
}

func (s *serviceImpl) AttendeeCentralStaffCheckin(userId int) *apperror.AppError {
	var checkin model.AttendeeCentralCheckin
	checkin.UserId = userId

	if err := s.repo.CreateCentralCheckin(&checkin); errors.Is(err, gorm.ErrDuplicatedKey) {
		return apperror.AlreadyCheckin
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return apperror.NotFound
	} else if err != nil {
		s.logger.Error("unable to central checkin", zap.Int("userId", userId), zap.Error(err))
		return apperror.InternalError
	}

	return nil
}
