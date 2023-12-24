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

	if container.Config.AppConfig.Env == "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/_hc", container.HcHandler.HealthCheck)
	r.GET("/featureflag/live", container.FeatureflagHandler.GetLivestreamInfo)
	r.POST("/register", container.RegisterHandler.CreateUser)
	r.GET("/register/:id", container.RegisterHandler.GetUserById)
	r.GET("/login", container.LoginHandler.GoogleLogin)
	r.GET("/login/callback", container.LoginHandler.GoogleCallback)
	r.GET("/live", container.FeatureflagHandler.GetLivestreamInfo)
	r.GET("/events", container.EventHandler.GetAllEvents)
	r.GET("/events/:eventId", container.EventHandler.GetEventById)

	if err := r.Run(fmt.Sprintf(":%v", container.Config.AppConfig.Port)); err != nil {
		container.Logger.Fatal("unable to start server")
	}
}
