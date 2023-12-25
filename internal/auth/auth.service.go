package auth

import (
	"context"
	"strings"

	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"golang.org/x/oauth2/google"
)

type Service interface {
	GoogleLogin() (string, *apperror.AppError)
	GoogleCallback(code string) (string, *apperror.AppError)
	Register (data *RegisterDTO) (*User, *apperror.AppError)
	GetUserFromJWTToken(tokenString string) (*User, *apperror.AppError)
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

func (s *serviceImpl) GoogleLogin() (string, *apperror.AppError) {
	url := s.oauthConfig.AuthCodeURL("state")
	if url == "" {
		s.logger.Error("Failed to get AuthCodeURL")
		return "", apperror.InternalError
	}
	return url , nil
}

func (s *serviceImpl) GoogleCallback(code string) (string, *apperror.AppError) {
	token, err := s.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		s.logger.Error("Failed to exchange code for token", zap.Error(err))
		return "", apperror.InternalError
	}

	idToken := token.Extra("id_token")
	if idToken == nil {
		s.logger.Error("ID token not found")
		return "", apperror.InternalError
	}


	return idToken.(string), nil
}

func (s *serviceImpl) Register(data *RegisterDTO) (*User, *apperror.AppError) {
	user, err := s.repo.GetUserByEmail(data.Email)
	dataUser := ConvertRegisterDTOToUser(data)

	if err == nil {
		user, err := s.repo.UpdateUser(user.ID, dataUser)
		if err != nil {
			s.logger.Error("Failed to update user", zap.Error(err))
			return nil, apperror.InternalError
		}
		return user, nil
	} else {
		user, err := s.repo.CreateUser(dataUser)
		if err != nil {
			s.logger.Error("Failed to create user", zap.Error(err))
			return nil, apperror.InternalError
		}
		return user, nil
}
}

func (s *serviceImpl) GetUserFromJWTToken(tokenString string) (*User, *apperror.AppError) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return nil, apperror.InvalidToken
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := idtoken.Validate(context.Background(), tokenString, s.cfg.OAuth2Config.ClientID)
	if err != nil {
		s.logger.Error("Failed to validate token", zap.Error(err))
		return nil, apperror.InvalidToken
	}

	email, ok := token.Claims["email"].(string)
	if !ok || email == "" {
		s.logger.Error("Email not found in token claims")
		return nil, apperror.InvalidToken
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		s.logger.Error("Could not retrieve user from database", zap.Error(err))
		return nil, apperror.InternalError
	}

	return user, nil
}


func ConvertRegisterDTOToUser(dto *RegisterDTO) *User {
	return &User{
		Gender:            dto.Gender,
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Email:             dto.Email,
		School:            dto.School,
		BirthDate:         dto.BirthDate,
		Address:           dto.Address,
		FromAbroad:        dto.FromAbroad,
		Allergy:           dto.Allergy,
		MedicalCondition:  dto.MedicalCondition,
		JoinCUReason:      dto.JoinCUReason,
		NewsSource:        dto.NewsSource,
		Status:            dto.Status,
		Grade:             dto.Grade,
		DesiredRounds:     dto.DesiredRounds,
		InterestedFaculties: dto.InterestedFaculties,
	}
}