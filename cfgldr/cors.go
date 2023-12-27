package cfgldr

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsHandler gin.HandlerFunc

func makeCorsConfig(cfg *Config) gin.HandlerFunc {
	if cfg.AppConfig.IsDevelopment() {
		return cors.New(cors.Config{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOriginFunc: func(string) bool {
				return true
			},
		})

	}

	allowOrigins := strings.Split(cfg.CorsConfig.AllowOrigins, ",")

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func MakeCorsConfig(cfg *Config) CorsHandler {
	return CorsHandler(makeCorsConfig(cfg))
}
