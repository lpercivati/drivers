package services

import (
	"root/bodies"
	"root/models"
	services "root/services/interfaces"
	"time"
)

type DriverService struct {
	Repository services.DriverRepository
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

func (service *DriverService) GetDriverByEmail(email string) (models.Driver, error) {
	return service.Repository.GetDriverByEmail(email)
}

func (service *DriverService) GetAvailableDrivers() ([]models.Driver, error) {
	return service.Repository.GetAvailableDrivers()
}
