package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/utils"
	"go.uber.org/zap"
)

type Handler interface {
	SubmitFeedback(c *gin.Context)
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
// @summary Submit feedback form
// @description Submit feedback form
// @id SubmitFeedback
// @produce json
// @Security Bearer
// @tags feedback
// @router /feedback [post]
// @param user body dto.SubmitFeedbackDTO true "Feedback dto"
// @success 204
func (h *handlerImpl) SubmitFeedback(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	var body dto.SubmitFeedbackDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	if apperr := h.svc.SubmitFeedback(&body, email); apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
