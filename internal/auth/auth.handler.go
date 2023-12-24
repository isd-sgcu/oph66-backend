package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"go.uber.org/zap"
)

type Handler interface {
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
	Register(c *gin.Context)
	GetProfile(c *gin.Context)
}

type handlerImpl struct {
	svc Service
	cfg *cfgldr.Config
	logger *zap.Logger
}

func NewHandler(svc Service, cfg *cfgldr.Config, logger *zap.Logger) Handler {
	return &handlerImpl{
		svc: svc,
		cfg: cfg,
		logger: logger,
	}
}

func (h *handlerImpl) GoogleLogin(c *gin.Context) {
	oauthConfig := h.initializeOAuthConfig()

	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *handlerImpl) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	oauthConfig := h.initializeOAuthConfig()

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		h.logger.Error("Failed to exchange code for token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	idToken := token.Extra("id_token")
	if idToken == nil {
		h.logger.Error("ID token not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID token not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token.AccessToken,
		"id_token":     idToken.(string), 
	})
}

func (h *handlerImpl) initializeOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     h.cfg.OAuth2Config.ClientID,
		ClientSecret: h.cfg.OAuth2Config.ClientSecret,
		RedirectURL:  h.cfg.OAuth2Config.RedirectURL,
		Scopes: h.cfg.OAuth2Config.Scopes,
		Endpoint: google.Endpoint,
	}
}

func (h *handlerImpl) Register(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CreateUser(&user); err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user, "id": user.ID})
}

func (h *handlerImpl) GetProfile(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        h.logger.Error("Authorization header is missing")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
        return
    }

    email, err := h.extractEmailFromJWTToken(authHeader)
    if err != nil {
        h.logger.Error("Invalid token", zap.Error(err))
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // Return specific error message
        return
    }

    user, err := h.svc.GetUserByEmail(email)
	if err != nil {
		h.logger.Error("Failed to get user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	DesiredRounds, err := h.svc.GetDesiredRoundsByUserId(user.ID)
	if err != nil {
		h.logger.Error("Failed to get desired rounds", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get desired rounds"})
		return
	}

	InterestedFaculties, err := h.svc.GetInterestedFacultiesByUserId(user.ID)
	if err != nil {
		h.logger.Error("Failed to get interested faculties", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get interested faculties"})
		return
	}
	

	user.DesiredRounds = DesiredRounds
	user.InterestedFaculties = InterestedFaculties

	c.JSON(http.StatusOK, gin.H{"user": user, "id": user.ID})
}

func (h *handlerImpl) extractEmailFromJWTToken(tokenString string) (string, *apperror.AppError) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		h.logger.Error("could not parse token", zap.Error(err))
		return "", apperror.InvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		h.logger.Error("invalid token claims")
		return "", apperror.InvalidToken
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		h.logger.Error("email not found in token claims")
		return "", apperror.InvalidToken
	}

	return email, nil
}