package repository

import (
	"log"
	"root/config"
	"root/models"
)

type DriverRepository struct {
}

func (_ *DriverRepository) Create(driver models.Driver) (models.Driver, error) {

	statement, _ := config.DB.Prepare("INSERT INTO drivers (Fullname, Email, PasswordHash, IsAdmin, DateCreation) VALUES (?, ?, ?, ?, ?)")
	result, err := statement.Exec(driver.Fullname, driver.Email, driver.PasswordHash, driver.IsAdmin, driver.DateCreation)
	if err == nil {
		id, _ := result.LastInsertId()
		driver.Id = int(id)

		return driver, err
	}

	log.Println("Unable to create driver", err.Error())
	return models.Driver{}, err
}

func (_ *DriverRepository) GetDrivers(page int) ([]models.Driver, error) {
	//limit := 5
	//offset := limit * (page - 1)
	rows, err := config.DB.Query("SELECT * FROM drivers")
	drivers := []models.Driver{}

	if err == nil {
		for rows.Next() {
			var currentDriver models.Driver

			rows.Scan(
				&currentDriver.Id,
				&currentDriver.Fullname,
				&currentDriver.Email,
				&currentDriver.PasswordHash,
				&currentDriver.IsAdmin,
				&currentDriver.DateCreation,
			)
			drivers = append(drivers, currentDriver)
		}
		return drivers, err
	}

	log.Println("Get driver fail", err.Error())
	return drivers, err
}
