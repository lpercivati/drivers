package services

import (
	"root/bodies"
	"root/models"
	"time"
)

type DriverRepository interface {
	Create(driver models.Driver) (models.Driver, error)
	GetDrivers(page int) ([]models.Driver, error)
}

type DriverService struct {
	Repository DriverRepository
}

func (service *DriverService) Create(data bodies.DriverBody) (models.Driver, error) {

	var driver = models.Driver{
		Fullname:     data.Fullname,
		Email:        data.Email,
		PasswordHash: data.Password,
		IsAdmin:      data.IsAdmin,
		DateCreation: time.Now().UTC(),
	}

	return service.Repository.Create(driver)
}

func (service *DriverService) GetDrivers(page int) ([]models.Driver, error) {
	return service.Repository.GetDrivers(page)
}
