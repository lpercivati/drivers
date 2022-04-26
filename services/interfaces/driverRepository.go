package services

import "root/models"

type DriverRepository interface {
	Create(driver models.Driver) (models.Driver, error)
	GetDrivers(page int, count int) ([]models.Driver, error)
	GetDriverByEmail(email string) (models.Driver, error)
	GetAvailableDrivers() ([]models.Driver, error)
}
