package main

import (
	"log"
	"root/config"
	"root/di"
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
		var driverController = di.GetDriverController(driverService)

		driverServer.GET("/drivers/:id", driverController.Get)
		driverServer.POST("/drivers", driverController.Add)

		driverServer.Run()

	}
}
