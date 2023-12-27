package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
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

// GoogleLogin godoc
// @summary Redirect to Google login page
// @description Redirect to Google login page
// @id GoogleLogin
// @produce json
// @tags auth
// @Security Bearer
// @router /auth/login [get]
func (h *handlerImpl) GoogleLogin(c *gin.Context) {
	url := h.svc.GoogleLogin()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback godoc
// @summary receive a token after successfully login with Google
// @description After successfully logging in with a @chula account, you'll receive a token. If you attempt to log in using a different domain, Google will not allow the login
// @id GoogleCallback
// @produce json
// @tags auth
// @Security Bearer
// @router /auth/callback [get]
// @success 200 {object} auth.CallbackResponse
// @Failure 500 {object} auth.CallbackErrorResponse
// @Failure 404 {object} auth.CallbackInvalidResponse
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

// Register godoc
// @summary Register
// @description Register new account with email
// @id Register
// @produce json
// @tags auth
// @Security Bearer
// @router /auth/register [post]
// @param user body auth.MockUser true "User"
// @success 200 {object} auth.MockRegisterResponse
// @Failure 500 {object} auth.RegisterErrorResponse
// @Failure 404 {object} auth.RegisterInvalidResponse
// @Failure 401 {object} auth.RegisterUnauthorized
// @Failure 498 {object} auth.RegisterInvalidToken
func (h *handlerImpl) Register(c *gin.Context) {
	var data RegisterRequestDTO
	var user model.User
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		utils.ReturnError(c, apperror.InvalidToken)
		return
	}
	token := strings.Replace(authHeader, "Bearer ", "", 1)

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	apperr := h.svc.Register(c, &data, token, &user)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	response := RegisterResponse{
		User: &user,
	}

	c.JSON(http.StatusOK, response)
}

// GetProfile godoc
// @summary Get Profile of current user
// @description Get Profile of current user
// @id GetProfile
// @produce json
// @tags auth
// @Security Bearer
// @router /auth/me [get]
// @success 200 {object} auth.MockGetProfileResponse
// @Failure 500 {object} auth.GetProfileErrorResponse
// @Failure 401 {object} auth.GetProfileUnauthorized
// @Failure 404 {object} auth.GetProfileUserNotFound
func (h *handlerImpl) GetProfile(c *gin.Context) {
	var user model.User
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		utils.ReturnError(c, apperror.InvalidToken)
		return
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)

	apperr := h.svc.GetUserFromJWTToken(c, token, &user)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	response := GetProfileResponse{
		User: &user,
	}

	c.JSON(http.StatusOK, response)
}
