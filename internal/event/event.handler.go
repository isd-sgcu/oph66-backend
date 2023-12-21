package event

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
)

type Handler interface {
	GetAllEvents(c *gin.Context)
	GetEventById(c *gin.Context)
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
	hit, result, apperr := h.service.GetEventCache(context.Background(), "get_all_events")
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	} else if hit {
		c.String(http.StatusOK, result)
		return
	}

	events, apperr := h.service.GetAllEvents(c.Request.Context())
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	eventsJson, err := json.Marshal(events)
	if err != nil {
		utils.ReturnError(c, apperror.InternalError)
		return
	}

	apperr = h.service.SetEventCache(context.Background(), "get_all_events", string(eventsJson), time.Hour*6)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	c.String(http.StatusOK, string(eventsJson))
}

func (h *handlerImpl) GetEventById(c *gin.Context) {
	eventId := c.Param("eventId")

	hit, result, apperr := h.service.GetEventCache(context.Background(), eventId)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	} else if hit {
		c.String(http.StatusOK, result)
		return
	}

	event, apperr := h.service.GetEventById(c.Request.Context(), eventId)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		utils.ReturnError(c, apperror.InternalError)
		return
	}

	apperr = h.service.SetEventCache(context.Background(), eventId, string(eventJson), time.Hour*6)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	c.String(http.StatusOK, string(eventJson))
}
