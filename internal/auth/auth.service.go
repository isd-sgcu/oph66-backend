package auth

import (
	"context"

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
		return "", apperror.ServiceUnavailable
	}

	return rawIdToken.(string), nil
}

func (s *serviceImpl) Register(ctx context.Context, data *RegisterRequestDTO, token string, user *User) (apperr *apperror.AppError) {
	email, apperr := getEmailFromToken(ctx, token, s.cfg.OAuth2Config.ClientID)
	if apperr != nil {
		return apperr
	}

	err := s.repo.GetUserByEmail(user, email)
	if err != nil {
		user = ConvertRegisterRequestDTOToUser(data, email)
		err = s.repo.CreateUser(user)
		if err != nil {
			s.logger.Error("Failed to create user", zap.Error(err))
			return apperror.InternalError
		}
	} else {
		s.logger.Error("User already exists")
		return apperror.DuplicateEmail
	}

	return nil
}

func (s *serviceImpl) GetUserFromJWTToken(ctx context.Context, token string, user *User) (apperr *apperror.AppError) {
	email, apperr := getEmailFromToken(ctx, token, s.cfg.OAuth2Config.ClientID)
	if apperr != nil {
		return apperr
	}

	err := s.repo.GetUserByEmail(user, email)
	if err != nil {
		s.logger.Error("Failed to get user by email", zap.Error(err))
		return apperror.UserNotFound
	}

	return nil
}

func getEmailFromToken(ctx context.Context, tokenString string, clientID string) (email string, appErr *apperror.AppError) {
	token, err := idtoken.Validate(ctx, tokenString, clientID)
	if err != nil {
		return "", apperror.InvalidToken
	}

	email, ok := token.Claims["email"].(string)
	if !ok || email == "" {
		return "", apperror.InvalidToken
	}

	return email, nil
}
