package evtreg

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
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
// @param scheduleId path int true "schedule id"
// @param registerEventDto body dto.EventRegistrationDTO true "Event register body"
// @router /schedules/{scheduleId}/register [post]
func (h *handlerImpl) RegisterEvent(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	var body dto.EventRegistrationDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	scheduleIdstr := c.Param("scheduleId")
	scheduleId, err := strconv.Atoi(scheduleIdstr)
	if err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	if apperr := h.svc.RegisterEvent(c, email, scheduleId, &body); apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
