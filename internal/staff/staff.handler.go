package staff

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
	"go.uber.org/zap"
)

type Handler interface {
	AttendeeStaffCheckin(c *gin.Context)
}

func NewHandler(service Service, logger *zap.Logger) Handler {
	return &handlerImpl{
		service,
		logger,
	}
}

type handlerImpl struct {
	service Service
	logger  *zap.Logger
}

// GetAllEvents godoc
// @summary checkin attendee
// @description	Checkin attendee which perform by staff
// @id AttendeeStaffCheckin
// @produce	json
// @tags staff
// @router /staff/checkin/{userId} [post]
// @param userId path int true "User id"
// @Security Bearer
// @success	204
// @Failure	403	{object} dto.EventInvalidResponse
// @Failure	404	{object} dto.EventInvalidResponse
// @Failure	409	{object} dto.EventInvalidResponse
// @Failure	500	{object} dto.EventAllErrorResponse
func (h *handlerImpl) AttendeeStaffCheckin(c *gin.Context) {
	if c.GetString("role") != "central-staff" && c.GetString("role") != "faculty-staff" {
		utils.ReturnError(c, apperror.Forbidden)
		return
	}

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.ReturnError(c, apperror.BadRequest)
		return
	}

	var (
		apperr *apperror.AppError
	)

	switch c.GetString("role") {
	case "faculty-staff":
		faculty := c.GetString("faculty")
		department := c.GetString("department")
		apperr = h.service.AttendeeFacultyStaffCheckin(userId, department, faculty)
	case "central-staff":
		apperr = h.service.AttendeeCentralStaffCheckin(userId)
	}

	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
