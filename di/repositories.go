package di

import (
	"root/config"
	"root/repositories"
)

func GetDriverRepository() repositories.DriverRepository {
	return repositories.DriverRepository{
		DataBase: config.DB,
	}
}
