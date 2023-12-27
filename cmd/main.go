package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/di"
)

func main() {
	container, err := di.Init()
	if err != nil {
		panic(fmt.Sprintf("unable to init di: %v", err))
	}
	container.Logger.Info("init container successfully")

	if !container.Config.AppConfig.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(gin.HandlerFunc(container.CorsHandler))

	r.GET("/_hc", container.HcHandler.HealthCheck)
	r.GET("/live", container.FeatureflagHandler.GetLivestreamInfo)
	r.GET("/events", container.EventHandler.GetAllEvents)
	r.GET("/events/:eventId", container.EventHandler.GetEventById)
	r.POST("/auth/register", container.AuthHandler.Register)
	r.GET("/auth/me", container.AuthHandler.GetProfile)
	r.GET("/auth/login", container.AuthHandler.GoogleLogin)
	r.GET("/auth/callback", container.AuthHandler.GoogleCallback)

	if err := r.Run(fmt.Sprintf(":%v", container.Config.AppConfig.Port)); err != nil {
		container.Logger.Fatal("unable to start server")
	}
}
