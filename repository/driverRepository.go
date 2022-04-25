package repository

import (
	"log"
	"root/config"
	"root/models"
	"time"
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
	limit := 5
	offset := limit * (page - 1)

	rows, err := config.DB.Query(`
	SELECT * FROM drivers
	ORDER BY id
	LIMIT $1
	OFFSET $2
	`, limit, offset)
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

func (_ *DriverRepository) GetDriverByEmail(email string) (models.Driver, error) {
	rows, err := config.DB.Query("SELECT * FROM drivers WHERE email = $1", email)
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

		if len(drivers) == 0 {
			return models.Driver{}, nil
		}

		return drivers[0], err
	}

	log.Println("Get driver fail", err.Error())
	return models.Driver{}, err
}

func (_ *DriverRepository) GetAvailableDrivers() ([]models.Driver, error) {
	dateNow := time.Now().UTC().Format("2006-01-02")
	log.Println(dateNow)

	rows, err := config.DB.Query(`
	SELECT * FROM drivers d
		WHERE d.Id NOT IN (
			SELECT t.DriverId FROM trips t
				WHERE t.DateStart <= $1 AND t.DateEnd >= $1
		)
	`, dateNow)
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
	return []models.Driver{}, err
}
