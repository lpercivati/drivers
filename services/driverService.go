package services

import (
	"log"
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
	log.Println(email)
	driver, err := service.Repository.GetDriverByEmail(email)
	log.Println(driver.PasswordHash)
	return driver, err
}
