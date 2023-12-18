package featureflag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/utils"
)

type Handler interface {
	LivestreamEnabled(c *gin.Context)
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc,
	}
}

type handlerImpl struct {
	svc Service
}

func (h *handlerImpl) LivestreamEnabled(c *gin.Context) {
	flag, err := h.svc.GetFlag(c.Request.Context(), "livestream")
	if err != nil {
		utils.ReturnError(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"livestream_enabled": flag,
		},
	)
}
