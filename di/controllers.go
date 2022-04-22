package di

import (
	"root/controllers"
	"root/services"
)

func GetDriverController(service services.DriverService) controllers.DriverController {
	return controllers.DriverController{
		Service: &service,
	}
}
