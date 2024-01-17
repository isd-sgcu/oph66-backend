package feedback

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
	SubmitFeedback(dto *dto.SubmitFeedbackDTO, userEmail string) *apperror.AppError
}

func NewService(repo Repository, userRepo auth.Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo,
		userRepo,
		logger,
	}
}

type serviceImpl struct {
	repo     Repository
	userRepo auth.Repository
	logger   *zap.Logger
}

func (s *serviceImpl) SubmitFeedback(dto *dto.SubmitFeedbackDTO, userEmail string) *apperror.AppError {
	var user model.User

	if err := s.userRepo.GetUserByEmail(&user, userEmail); errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.NotFound
	} else if err != nil {
		s.logger.Error("unable to get user by email", zap.String("email", userEmail), zap.Any("submitFeedbackDto", dto))
		return apperror.InternalError
	}

	feedback := FeedbackDTOToModel(dto)
	feedback.UserId = user.Id

	if err := s.repo.CreateFeedback(&feedback); errors.Is(err, gorm.ErrDuplicatedKey) {
		return apperror.AlreadySubmitted
	} else if err != nil {
		s.logger.Error("unable to submit feedback form", zap.String("email", userEmail), zap.Any("submitFeedbackDto", dto), zap.Any("feedback", feedback))
		return apperror.InternalError
	}

	return nil
}
