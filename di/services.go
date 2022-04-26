package di

import (
	"root/repositories"
	"root/services"
)

func GetDriverService(repository repositories.DriverRepository) services.DriverService {
	return services.DriverService{
		Repository: &repository,
	}
}

func GetAuthService() services.AuthService {
	return services.AuthService{}
}
