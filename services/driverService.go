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

	driverResponse, err := service.Repository.Create(driver)

	if err != nil {
		log.Println("Error in driver creation")

		return models.Driver{}, err
	}

	return driverResponse, err
}

func (service *DriverService) GetDrivers(page int, count int) ([]models.Driver, error) {
	drivers, err := service.Repository.GetDrivers(page, count)

	if err != nil {
		log.Println("Error getting drivers")

		return []models.Driver{}, err
	}

	return drivers, err
}

func (service *DriverService) GetDriverByEmail(email string) (models.Driver, error) {
	driverResponse, err := service.Repository.GetDriverByEmail(email)

	if err != nil {
		log.Println("Error getting driver")

		return models.Driver{}, err
	}

	return driverResponse, err
}

func (service *DriverService) GetAvailableDrivers() ([]models.Driver, error) {
	drivers, err := service.Repository.GetAvailableDrivers()

	if err != nil {
		log.Println("Error getting drivers")

		return []models.Driver{}, err
	}

	return drivers, err
}
