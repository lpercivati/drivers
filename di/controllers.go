package di

import (
	"root/controllers"
	"root/services"
)

func GetDriverController(service services.DriverService, authService services.AuthService) controllers.DriverController {
	return controllers.DriverController{
		Service:     &service,
		AuthService: &authService,
	}
}

func GetAuthController(authService services.AuthService, driverService services.DriverService) controllers.AuthController {
	return controllers.AuthController{
		AuthService:   &authService,
		DriverService: &driverService,
	}
}
