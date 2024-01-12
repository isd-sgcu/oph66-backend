package staff

import (
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/auth"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	AttendeeFacultyStaffCheckin(userId int, departmentCode string, facultyCode string) (ciu *dto.AttendeeStaffCheckinUser, alreadyCheckin bool, apperr *apperror.AppError)
	AttendeeCentralStaffCheckin(userId int) (ciu *dto.AttendeeStaffCheckinUser, alreadyCheckin bool, apperr *apperror.AppError)
}

func NewService(repo Repository, authRepo auth.Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		authRepo,
		logger,
	}
}

type serviceImpl struct {
	repo     Repository
	authRepo auth.Repository
	logger   *zap.Logger
}

func (s *serviceImpl) AttendeeFacultyStaffCheckin(userId int, departmentCode string, facultyCode string) (*dto.AttendeeStaffCheckinUser, bool, *apperror.AppError) {
	var checkin model.AttendeeFacultyCheckin
	checkin.UserId = userId
	checkin.DepartmentCode = departmentCode
	checkin.FacultyCode = facultyCode

	err := s.repo.CreateFacultyCheckin(&checkin)
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, false, apperror.NotFound
	} else if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
		s.logger.Error("unable to faculty checkin", zap.Int("userId", userId), zap.String("departmentCode", departmentCode), zap.String("facultyCode", facultyCode), zap.Error(err))
		return nil, false, apperror.InternalError
	}

	alreadyCheckin := errors.Is(err, gorm.ErrDuplicatedKey)

	var user model.User
	if err := s.authRepo.GetUserById(&user, userId); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, apperror.UserNotFound
	} else if err != nil {
		s.logger.Error("unable to find user after faculty checkin", zap.Int("userId", userId), zap.String("departmentCode", departmentCode), zap.String("facultyCode", facultyCode), zap.Error(err))
		return nil, false, apperror.InternalError
	}

	ciu := UserModelToCheckinUser(&user)

	return &ciu, alreadyCheckin, nil
}

func (s *serviceImpl) AttendeeCentralStaffCheckin(userId int) (*dto.AttendeeStaffCheckinUser, bool, *apperror.AppError) {
	var checkin model.AttendeeCentralCheckin
	checkin.UserId = userId

	err := s.repo.CreateCentralCheckin(&checkin)
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, false, apperror.NotFound
	} else if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
		s.logger.Error("unable to central checkin", zap.Int("userId", userId), zap.Error(err))
		return nil, false, apperror.InternalError
	}

	alreadyCheckin := errors.Is(err, gorm.ErrDuplicatedKey)

	var user model.User
	if err := s.authRepo.GetUserById(&user, userId); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, apperror.UserNotFound
	} else if err != nil {
		s.logger.Error("unable to find user after central checkin", zap.Int("userId", userId), zap.Error(err))
		return nil, false, apperror.InternalError
	}

	ciu := UserModelToCheckinUser(&user)

	return &ciu, alreadyCheckin, nil
}
