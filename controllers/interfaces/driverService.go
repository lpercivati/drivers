package controllers

import (
	"root/bodies"
	"root/models"
)

type DriverService interface {
	Create(data bodies.DriverBody) (models.Driver, error)
	GetDrivers(page int, count int) ([]models.Driver, error)
	GetDriverByEmail(email string) (models.Driver, error)
	GetAvailableDrivers() ([]models.Driver, error)
}
