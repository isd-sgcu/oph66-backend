package evtreg

import (
	"net/http"
	"strconv"

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
// @param scheduleId path int true "schedule id"
// @router /schedules/{scheduleId}/register [post]
func (h *handlerImpl) RegisterEvent(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		utils.ReturnError(c, apperror.Unauthorized)
		return
	}

	scheduleIdstr := c.Param("scheduleId")
	scheduleId, err := strconv.Atoi(scheduleIdstr)
	if err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	if apperr := h.svc.RegisterEvent(email, scheduleId); apperr != nil {
		utils.ReturnError(c, apperr)
	}

	c.AbortWithStatus(http.StatusNoContent)
}
