package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
)

func ReturnError(c *gin.Context, err *apperror.AppError) {
	c.JSON(
		err.HttpCode,
		gin.H{
			"instance": c.Request.URL.Path,
			"title":    err.Id,
		},
	)
}
