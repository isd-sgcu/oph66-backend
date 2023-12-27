package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	HealthCheck(c *gin.Context)
}

func NewHandler() Handler {
	return &handlerImpl{}
}

type handlerImpl struct {
}

// HealthCheck godoc
// @summary Health Check
// @description Health Check for the service
// @id HealthCheck
// @produce plain
// @tags healthcheck
// @Security Bearer
// @router /_hc [get]
// @Success 200 {string} string "OK"
func (h *handlerImpl) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
