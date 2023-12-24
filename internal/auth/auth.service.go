package auth

import (
	"github.com/isd-sgcu/oph66-backend/apperror"
	"go.uber.org/zap"
)

type Service interface {
	CreateUser(user *User) *apperror.AppError
	GetUserByEmail(email string) (*User, *apperror.AppError)
	GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, *apperror.AppError)
	GetDesiredRoundsByUserId(id uint) ([]DesiredRounds, *apperror.AppError)
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		logger,
	}
}


type serviceImpl struct {
	repo Repository
	logger *zap.Logger
}

func (s *serviceImpl) CreateUser(user *User) *apperror.AppError {
	err := s.repo.Create(user)
	if err!= nil {	
		s.logger.Error("could not create user in database", zap.Error(err))
		return apperror.InternalError
	}
	return nil
}

func (s *serviceImpl) GetUserByEmail(email string) (*User, *apperror.AppError) {
	user , err := s.repo.GetByEmail(email)
	if err != nil {
		s.logger.Error("could not retrieve user from database", zap.Error(err))
		return nil, apperror.InternalError
	}
	return user, nil
}

func (s *serviceImpl) GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, *apperror.AppError) {
	user , err := s.repo.GetInterestedFacultiesByUserId(id)
	if err != nil {
		s.logger.Error("could not retrieve interested faculties from database", zap.Error(err))
		return nil, apperror.InternalError
	}
	return user, nil
}

func (s *serviceImpl) GetDesiredRoundsByUserId(id uint) ([]DesiredRounds, *apperror.AppError) {
	user , err := s.repo.GetDesiredRoundsByUserId(id)
	if err != nil {
		s.logger.Error("could not retrieve desired rounds from database", zap.Error(err))
		return nil, apperror.InternalError
	}
	return user, nil
}
