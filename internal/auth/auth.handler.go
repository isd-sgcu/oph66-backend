package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
	"go.uber.org/zap"
)

type Handler interface {
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
	Register(c *gin.Context)
	GetProfile(c *gin.Context)
}

type handlerImpl struct {
	svc    Service
	logger *zap.Logger
}

func NewHandler(svc Service, logger *zap.Logger) Handler {
	return &handlerImpl{
		svc,
		logger,
	}
}

func (h *handlerImpl) GoogleLogin(c *gin.Context) {
	url, apperr := h.svc.GoogleLogin()
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
	c.JSON(http.StatusOK, gin.H{"message": "GoogleLogin successful"})
}

func (h *handlerImpl) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token , apperr := h.svc.GoogleCallback(code)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "GoogleCallback successful", "token": token})
}

func (h *handlerImpl) Register(c *gin.Context) {
	var data RegisterDTO

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	user, apperr := h.svc.Register(&data)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}


func (h *handlerImpl) GetProfile(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	user, apperr := h.svc.GetUserFromJWTToken(authHeader) 
	if apperr != nil {
		utils.ReturnError(c, apperr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user from JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "id": user.ID})
}