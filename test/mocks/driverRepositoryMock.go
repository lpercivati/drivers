package mocks

import (
	"errors"
	"root/models"
)

type DriverRepository struct {
}

func (service *DriverRepository) GetDrivers(page int, count int) ([]models.Driver, error) {
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

func (service *DriverRepository) Create(data models.Driver) (models.Driver, error) {
	return models.Driver{}, errors.New("error")
}
func (service *DriverRepository) GetDriverByEmail(email string) (models.Driver, error) {
	return models.Driver{}, errors.New("error")
}
func (service *DriverRepository) GetAvailableDrivers() ([]models.Driver, error) {
	return []models.Driver{}, errors.New("error")
}
