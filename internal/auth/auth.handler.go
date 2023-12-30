package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
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
// @param code query string true "Authorization code"
// @router /auth/callback [get]
// @success 200 {object} dto.CallbackResponse
// @Failure 500 {object} dto.CallbackErrorResponse
// @Failure 404 {object} dto.CallbackInvalidResponse
func (h *handlerImpl) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, apperr := h.svc.GoogleCallback(c, code)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	response := dto.GoogleCallbackResponse{
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
// @param user body dto.RegisterRequestDTO true "User"
// @success 200 {object} dto.RegisterResponse
// @Failure 500 {object} dto.RegisterErrorResponse
// @Failure 404 {object} dto.RegisterInvalidResponse
// @Failure 401 {object} dto.RegisterUnauthorized
// @Failure 498 {object} dto.RegisterInvalidToken
func (h *handlerImpl) Register(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	var data dto.RegisterRequestDTO

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	user, apperr := h.svc.Register(email, &data)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}
	response := dto.RegisterResponse{
		User: user,
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
// @success 200 {object} dto.GetProfileResponse
// @Failure 500 {object} dto.GetProfileErrorResponse
// @Failure 401 {object} dto.GetProfileUnauthorized
// @Failure 404 {object} dto.GetProfileUserNotFound
func (h *handlerImpl) GetProfile(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	user, apperr := h.svc.GetUserFromJWTToken(email)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	response := dto.GetProfileResponse{
		User: user,
	}

	c.JSON(http.StatusOK, response)
}
