package main

import (
	"fmt"

	"github.com/isd-sgcu/oph66-backend/di"
	"github.com/isd-sgcu/oph66-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                      OPH-66 Backend API
// @version                    1.0
// @description                Documentation outlines the specifications and endpoints for the OPH-66 Backend API.
// @Schemes                    http https
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.
func main() {
	container, err := di.Init()
	if err != nil {
		panic(fmt.Sprintf("unable to init di: %v", err))
	}
	container.Logger.Info("init container successfully")

	docs.SwaggerInfo.Host = container.Config.AppConfig.Host

	r := container.Router

	r.GET("/_hc", container.HcHandler.HealthCheck)
	r.GET("/live", container.FeatureflagHandler.GetLivestreamInfo)
	r.GET("/events", container.EventHandler.GetAllEvents)
	r.GET("/events/:eventId", container.EventHandler.GetEventById)
	r.POST("/schedules/:scheduleId/register", container.EvtregHandler.RegisterEvent)
	r.POST("/staff/checkin/:userId", container.StaffHandler.AttendeeStaffCheckin)
	r.POST("/auth/register", container.AuthHandler.Register)
	r.GET("/auth/me", container.AuthHandler.GetProfile)
	r.GET("/auth/login", container.AuthHandler.GoogleLogin)
	r.GET("/auth/callback", container.AuthHandler.GoogleCallback)

	if container.Config.AppConfig.IsDevelopment() {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	if err := r.Run(fmt.Sprintf(":%v", container.Config.AppConfig.Port)); err != nil {
		container.Logger.Fatal("unable to start server")
	}
}
