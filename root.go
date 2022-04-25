package main

import (
	"log"
	"root/config"
	"root/di"
	"root/middlewares"
	"root/migrations"

	"github.com/gin-gonic/gin"
)

func main() {

	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Run all migrations
		migrations.Run()

		driverServer := gin.Default()

		var driverRepository = di.GetDriverRepository()
		var driverService = di.GetDriverService(driverRepository)
		var authService = di.GetAuthService()

		var driverController = di.GetDriverController(driverService, authService)
		var authController = di.GetAuthController(authService, driverService)

		api := driverServer.Group("/api")
		{
			api.POST("/token", authController.GenerateToken)
			api.POST("/drivers", driverController.Add)
			secured := api.Group("/secured").Use(middlewares.Auth())
			{
				secured.GET("/drivers", driverController.Get)

			}
		}

		driverServer.Run()

	}
}
