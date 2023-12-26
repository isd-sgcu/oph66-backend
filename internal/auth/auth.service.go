package auth

import (
	"context"
	"strings"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type Service interface {
	GoogleLogin() (url string)
	GoogleCallback(ctx context.Context, code string) (idToken string, appErr *apperror.AppError)
	Register(ctx context.Context, data *RegisterRequestDTO, tokenString string, user *User) (appErr *apperror.AppError)
	GetUserFromJWTToken(ctx context.Context, tokenString string, user *User) (appErr *apperror.AppError)
}

func NewService(repo Repository, logger *zap.Logger, cfg *cfgldr.Config) Service {
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.OAuth2Config.ClientID,
		ClientSecret: cfg.OAuth2Config.ClientSecret,
		RedirectURL:  cfg.OAuth2Config.RedirectURL,
		Scopes:       cfg.OAuth2Config.Scopes,
		Endpoint:     google.Endpoint,
	}
	return &serviceImpl{
		repo,
		cfg,
		logger,
		oauthConfig,
	}
}

type serviceImpl struct {
	repo        Repository
	cfg         *cfgldr.Config
	logger      *zap.Logger
	oauthConfig *oauth2.Config
}

func (s *serviceImpl) GoogleLogin() (url string) {
	url = s.oauthConfig.AuthCodeURL("state")
	return url
}

func (s *serviceImpl) GoogleCallback(ctx context.Context, code string) (idToken string, appErr *apperror.AppError) {
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		s.logger.Error("Failed to exchange code for token", zap.Error(err))
		return "", apperror.InternalError
	}

	rawIdToken := token.Extra("id_token")
	if rawIdToken == nil {
		s.logger.Error("ID token not found")
		return "", apperror.InternalError
	}

	return rawIdToken.(string), nil
}

func (s *serviceImpl) Register(ctx context.Context, data *RegisterRequestDTO, tokenString string, user *User) (appErr *apperror.AppError) {
	email, appErr := getEmailFromToken(ctx, tokenString, s.cfg.OAuth2Config.ClientID)
	if appErr != nil {
		return appErr
	}

	err := s.repo.GetUserByEmail(user, email)
	ConvertRegisterRequestDTOToUser(data, email, user.ID, user)

	if err == nil {
		err = s.repo.UpdateUser(user)
		if err != nil {
			s.logger.Error("Failed to update user", zap.Error(err))
			return apperror.InternalError
		}
	} else {
		err = s.repo.CreateUser(user)
		if err != nil {
			s.logger.Error("Failed to create user", zap.Error(err))
			return apperror.InternalError
		}
	}

	return nil
}

func (s *serviceImpl) GetUserFromJWTToken(ctx context.Context, tokenString string, user *User) (appErr *apperror.AppError) {
	email, appErr := getEmailFromToken(ctx, tokenString, s.cfg.OAuth2Config.ClientID)
	if appErr != nil {
		return appErr
	}

	err := s.repo.GetUserByEmail(user, email)
	if err != nil {
		s.logger.Error("Failed to get user by email", zap.Error(err))
		return apperror.InternalError
	}

	return nil
}

func getEmailFromToken(ctx context.Context, tokenString string, ClientID string) (email string, appErr *apperror.AppError) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", apperror.InvalidToken
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := idtoken.Validate(ctx, tokenString, ClientID)
	if err != nil {
		return "", apperror.InvalidToken
	}

	email, ok := token.Claims["email"].(string)
	if !ok || email == "" {
		return "", apperror.InvalidToken
	}

	return email, nil
}

func ConvertRegisterRequestDTOToUser(dto *RegisterRequestDTO, email string, id uint, user *User) {
	user.ID = id
	user.Gender = dto.Gender
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Email = email
	user.School = dto.School
	user.BirthDate = dto.BirthDate
	user.Address = dto.Address
	user.FromAbroad = dto.FromAbroad
	user.Allergy = dto.Allergy
	user.MedicalCondition = dto.MedicalCondition
	user.JoinCUReason = dto.JoinCUReason
	user.NewsSource = dto.NewsSource
	user.Status = dto.Status
	user.Grade = dto.Grade
	user.DesiredRounds = make([]DesiredRound, len(dto.DesiredRounds))
	user.InterestedFaculties = make([]InterestedFaculty, len(dto.InterestedFaculties))
	for i, desiredRound := range dto.DesiredRounds {
		user.DesiredRounds[i] = ConvertDesiredInfoToDesiredRound(&desiredRound, user)
	}
	for i, interestedFaculty := range dto.InterestedFaculties {
		user.InterestedFaculties[i] = ConvertFacultyInfoToInterestedFaculty(&interestedFaculty, user)
	}
}

func ConvertDesiredInfoToDesiredRound(dto *DesiredInfo, user *User) (desiredRound DesiredRound) {
	desiredRound.Order = dto.Order
	desiredRound.RoundCode = dto.Code
	return desiredRound
}

func ConvertFacultyInfoToInterestedFaculty(dto *FacultyInfo, user *User) (interestedFaculty InterestedFaculty) {
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
	interestedFaculty.UserID = user.ID
	return interestedFaculty
}
