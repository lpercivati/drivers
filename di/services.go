package di

import (
	"root/repository"
	"root/services"
)

func GetDriverService(repository repository.DriverRepository) services.DriverService {
	return services.DriverService{
		Repository: &repository,
	}
}

func GetAuthService() services.AuthService {
	return services.AuthService{}
}
