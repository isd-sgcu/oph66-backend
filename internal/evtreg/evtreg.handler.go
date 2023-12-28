package evtreg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
)

type Handler interface {
	RegisterEvent(c *gin.Context)
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc,
	}
}

var _ Handler = &handlerImpl{}

type handlerImpl struct {
	svc Service
}

// GoogleLogin godoc
// @summary Register event
// @description Register event
// @id RegisterEvent
// @produce json
// @tags event
// @Security Bearer
// @router /events/{eventId}/register [post]
func (h *handlerImpl) RegisterEvent(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	var body EventRegistrationDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	if apperr := h.svc.RegisterEvent(email, body.ScheduleId); apperr != nil {
		utils.ReturnError(c, apperr)
	}

	c.AbortWithStatus(http.StatusNoContent)
}
