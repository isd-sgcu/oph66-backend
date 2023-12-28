package featureflag

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/utils"
)

type Handler interface {
	GetLivestreamInfo(c *gin.Context)
}

func NewHandler(svc Service, cache Cache) Handler {
	return &handlerImpl{
		svc,
		cache,
	}
}

var _ Handler = &handlerImpl{}

type handlerImpl struct {
	svc   Service
	cache Cache
}

// GetLivestreamInfo godoc
// @summary Get livestream flag
// @description	Get livestream flag
// @id GetLivestreamInfo
// @produce json
// @tags FeatureFlag
// @router /live [get]
// @success 200 {object} dto.response
// @Failure 500 {object} dto.errorResponse
// @Failure 404 {object} dto.invalidResponse
func (h *handlerImpl) GetLivestreamInfo(c *gin.Context) {
	cacheKey := "livestream"

	cached, err := h.cache.Get(c, cacheKey)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	if cached != nil {
		c.JSON(http.StatusOK, cached)
		return
	} else {
		data, err := h.svc.GetFlag(c, cacheKey)
		if err != nil {
			utils.ReturnError(c, err)
			return
		}

		if err = h.cache.Set(c, cacheKey, data, time.Duration(data.CacheDuration)*time.Second); err != nil {
			utils.ReturnError(c, err)
			return
		}

		c.JSON(
			http.StatusOK,
			data,
		)
	}
}
