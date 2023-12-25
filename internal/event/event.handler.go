package event

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/utils"
	"go.uber.org/zap"
)

type Handler interface {
	GetAllEvents(c *gin.Context)
	GetEventById(c *gin.Context)
}

func NewHandler(service Service, cache Cache, logger *zap.Logger) Handler {
	return &handlerImpl{
		service,
		cache,
		logger,
	}
}

type handlerImpl struct {
	service Service
	cache   Cache
	logger  *zap.Logger
}

func (h *handlerImpl) GetAllEvents(c *gin.Context) {
	hit, result, apperr := h.cache.Get(c.Request.Context(), "get_all_events")
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	} else if hit {
		setHeader(c)
		c.String(http.StatusOK, result)
		return
	}

	events, apperr := h.service.GetAllEvents()
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	eventsJson, err := json.Marshal(events)
	if err != nil {
		h.logger.Error("could not serialize into json format", zap.Error(err))
		utils.ReturnError(c, apperror.InternalError)
		return
	}

	apperr = h.cache.Set(c.Request.Context(), "get_all_events", string(eventsJson), time.Hour*6)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	setHeader(c)
	c.String(http.StatusOK, string(eventsJson))
}

func (h *handlerImpl) GetEventById(c *gin.Context) {
	eventId := c.Param("eventId")

	hit, result, apperr := h.cache.Get(c.Request.Context(), fmt.Sprintf("get_event_by_id-%v", eventId))
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	} else if hit {
		setHeader(c)
		c.String(http.StatusOK, result)
		return
	}

	event, apperr := h.service.GetEventById(eventId)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		h.logger.Error("could not serialize into json format", zap.Error(err))
		utils.ReturnError(c, apperror.InternalError)
		return
	}

	apperr = h.cache.Set(c.Request.Context(), fmt.Sprintf("get_event_by_id-%v", eventId), string(eventJson), time.Hour*6)
	if apperr != nil {
		utils.ReturnError(c, apperr)
		return
	}

	setHeader(c)
	c.String(http.StatusOK, string(eventJson))
}

func setHeader(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=3600")
}