package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/isd-sgcu/oph66-backend/internal/middleware"
)

type Router struct {
	*gin.Engine
}

func NewRouter(config *cfgldr.Config, corsHandler cfgldr.CorsHandler, authMiddleware middleware.AuthMiddleware) *Router {

	if !config.AppConfig.IsDevelopment() {

		gin.SetMode(gin.ReleaseMode)

	}

	r := gin.Default()

	r.Use(gin.HandlerFunc(corsHandler))

	r.Use(gin.HandlerFunc(authMiddleware))

	return &Router{r}

}
