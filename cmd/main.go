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
	r.GET("/featureflag/live", container.FeatureflagHandler.LivestreamEnabled)
	r.GET("/event/all", container.EventHandler.GetAllEvents)

	// change this back to :PORT later
	if err := r.Run(fmt.Sprintf("localhost:%v", container.Config.AppConfig.Port)); err != nil {
		container.Logger.Fatal("unable to start server")
	}
}
