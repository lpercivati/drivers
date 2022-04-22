package di

import "root/repository"

func GetDriverRepository() repository.DriverRepository {
	return repository.DriverRepository{}
}
