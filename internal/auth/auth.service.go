package auth

import (
	"context"
	"errors"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type Service interface {
	GoogleLogin() (url string)
	GoogleCallback(ctx context.Context, code string) (idToken string, appErr *apperror.AppError)
	Register(user *User, email string, data *RegisterRequestDTO) *apperror.AppError
	GetUserFromJWTToken(user *User, email string) *apperror.AppError
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
		return "", apperror.ServiceUnavailable
	}

	return rawIdToken.(string), nil
}

func (s *serviceImpl) Register(user *User, email string, data *RegisterRequestDTO) *apperror.AppError {
	var mUser model.User
	ConvertRegisterRequestDTOToUser(&mUser, data, email)
	err := s.repo.CreateUser(&mUser)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return apperror.DuplicateEmail
	} else if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return apperror.InternalError
	}

	UserModelToUserDTO(user, &mUser)

	return nil
}

func (s *serviceImpl) GetUserFromJWTToken(user *User, email string) *apperror.AppError {
	var mUser model.User
	err := s.repo.GetUserByEmail(&mUser, email)
	if err != nil {
		return apperror.UserNotFound
	}

	UserModelToUserDTO(user, &mUser)

	return nil
}
