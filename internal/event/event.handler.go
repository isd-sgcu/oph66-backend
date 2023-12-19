package event

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
)

type Handler interface {
	GetAllEvents(c *gin.Context)
}

func NewHandler(service Service) Handler {
	return &handlerImpl{
		service,
	}
}

type handlerImpl struct {
	service Service
}

func (h *handlerImpl) GetAllEvents(c *gin.Context) {
	events, err := h.service.GetAllEvents(c.Request.Context())
	if err != nil {
		utils.ReturnError(c, err)
		return
	}

	eventsJson, jsonerr := json.Marshal(events)
	if jsonerr != nil {
		utils.ReturnError(c, apperror.InternalError)
		return
	}

	c.String(http.StatusOK, string(eventsJson))
}
