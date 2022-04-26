package mocks

import (
	"errors"
	"root/bodies"
	"root/models"
)

type DriverService struct {
}

func (service *DriverService) GetDrivers(page int, count int) ([]models.Driver, error) {
	if page == 1 && count == 1 {
		return []models.Driver{
			{
				Id:       10,
				Fullname: "Leandro",
			},
		}, nil
	}

	return []models.Driver{}, errors.New("error")
}

func (service *DriverService) Create(data bodies.DriverBody) (models.Driver, error) {
	return models.Driver{}, errors.New("error")
}
func (service *DriverService) GetDriverByEmail(email string) (models.Driver, error) {
	return models.Driver{}, errors.New("error")
}
func (service *DriverService) GetAvailableDrivers() ([]models.Driver, error) {
	return []models.Driver{}, errors.New("error")
}
