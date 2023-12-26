package auth

import (
	"net/http"
	"strings"

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
	url := h.svc.GoogleLogin()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *handlerImpl) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, apperr := h.svc.GoogleCallback(c, code)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	response := GoogleCallbackResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerImpl) Register(c *gin.Context) {
	var data RegisterRequestDTO
	var user User
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		utils.ReturnError(c, apperror.InvalidToken)
		return
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	apperr := h.svc.Register(c, &data, authHeader, &user)
	if apperr != nil {
		utils.ReturnError(c, apperr)
	}
	response := RegisterResponse{
		User: &user,
	}

	c.JSON(http.StatusOK, response)
}

func (h *handlerImpl) GetProfile(c *gin.Context) {
	var user User
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		utils.ReturnError(c, apperror.InvalidToken)
		return
	}

	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	apperr := h.svc.GetUserFromJWTToken(c, authHeader, &user)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	response := GetProfileResponse{
		User: &user,
	}

	c.JSON(http.StatusOK, response)
}
