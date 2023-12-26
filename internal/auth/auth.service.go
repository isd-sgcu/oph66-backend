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
	Register(ctx context.Context, data *RegisterDTO, tokenString string) (user *User, appErr *apperror.AppError)
	GetUserFromJWTToken(ctx context.Context, tokenString string) (user *User, appErr *apperror.AppError)
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

func (s *serviceImpl) Register(ctx context.Context, data *RegisterDTO, tokenString string) (user *User, appErr *apperror.AppError) {
	email, appErr := getEmailFromToken(ctx, s.logger, tokenString, s.cfg.OAuth2Config.ClientID)
	if appErr != nil {
		return nil, appErr
	}

	user, err := s.repo.GetUserByEmail(&User{}, email)

	dataUser := ConvertRegisterDTOToUser(data, email)

	if err == nil {
		user, err = s.repo.UpdateUser(user.ID, dataUser)
		if err != nil {
			s.logger.Error("Failed to update user", zap.Error(err))
			return nil, apperror.InternalError
		}
	} else {
		user, err = s.repo.CreateUser(dataUser)
		if err != nil {
			s.logger.Error("Failed to create user", zap.Error(err))
			return nil, apperror.InternalError
		}
	}

	return user, nil
}

func (s *serviceImpl) GetUserFromJWTToken(ctx context.Context, tokenString string) (user *User, appErr *apperror.AppError) {
	email, appErr := getEmailFromToken(ctx, s.logger, tokenString, s.cfg.OAuth2Config.ClientID)
	if appErr != nil {
		return nil, appErr
	}

	user, err := s.repo.GetUserByEmail(user, email)
	if err != nil {
		s.logger.Error("Failed to get user by email", zap.Error(err))
		return nil, apperror.InternalError
	}

	return user, nil
}

func getEmailFromToken(ctx context.Context, logger *zap.Logger, tokenString string, ClientID string) (email string, appErr *apperror.AppError) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", apperror.InvalidToken
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := idtoken.Validate(ctx, tokenString, ClientID)
	if err != nil {
		logger.Error("Failed to validate token", zap.Error(err))
		return "", apperror.InvalidToken
	}

	email, ok := token.Claims["email"].(string)
	if !ok || email == "" {
		logger.Error("Email not found in token claims")
		return "", apperror.InvalidToken
	}

	return email, nil
}

func ConvertRegisterDTOToUser(dto *RegisterDTO, email string) *User {
	return &User{
		Gender:              dto.Gender,
		FirstName:           dto.FirstName,
		LastName:            dto.LastName,
		Email:               email,
		School:              dto.School,
		BirthDate:           dto.BirthDate,
		Address:             dto.Address,
		FromAbroad:          dto.FromAbroad,
		Allergy:             dto.Allergy,
		MedicalCondition:    dto.MedicalCondition,
		JoinCUReason:        dto.JoinCUReason,
		NewsSource:          dto.NewsSource,
		Status:              dto.Status,
		Grade:               dto.Grade,
		DesiredRounds:       dto.DesiredRounds,
		InterestedFaculties: dto.InterestedFaculties,
	}
}