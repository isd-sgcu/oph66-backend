package event

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/dto"
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

// GetAllEvents godoc
// @summary Get all events
// @description	Get all events as array of events
// @id GetAllEvents
// @produce	json
// @tags event
// @router /events [get]
// @success	200	{object} dto.GetAllEventResponse
// @Failure	500	{object} dto.EventAllErrorResponse
// @Failure	404	{object} dto.EventInvalidResponse
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

	response := dto.GetAllEventResponse{
		Events: events,
	}

	eventsJson, err := json.Marshal(response)
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
	c.JSON(http.StatusOK, response)
}

// GetEvent godoc
// @summary get event by id
// @description Get event by id
// @id GetEventById
// @produce json
// @tags event
// @param eventId path string true "event id"
// @router /events/{eventId} [get]
// @success 200 {object} dto.GetEventByIdResponse
// @Failure 500 {object} dto.EventErrorResponse
// @Failure 404 {object} dto.EventInvalidResponse
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

	response := dto.GetEventByIdResponse{
		Event: event,
	}

	eventJson, err := json.Marshal(response)
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
	c.JSON(http.StatusOK, response)
}

func setHeader(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=3600")
}
