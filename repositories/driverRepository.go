package repositories

import (
	"log"
	"root/models"
	repositories "root/repositories/interfaces"
	"time"
)

type DriverRepository struct {
	DataBase repositories.DataBaseRepository
}

func (repository *DriverRepository) Create(driver models.Driver) (models.Driver, error) {

	statement, _ := repository.DataBase.Prepare("INSERT INTO drivers (Fullname, Email, PasswordHash, IsAdmin, DateCreation) VALUES (?, ?, ?, ?, ?)")
	result, err := statement.Exec(driver.Fullname, driver.Email, driver.PasswordHash, driver.IsAdmin, driver.DateCreation)
	if err == nil {
		id, _ := result.LastInsertId()
		driver.Id = int(id)

		return driver, err
	}

	log.Println("Unable to create driver", err.Error())
	return models.Driver{}, err
}

func (repository *DriverRepository) GetDrivers(page int, count int) ([]models.Driver, error) {
	limit := count
	offset := limit * (page - 1)

	rows, err := repository.DataBase.Query(`
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

func (repository *DriverRepository) GetDriverByEmail(email string) (models.Driver, error) {
	rows, err := repository.DataBase.Query("SELECT * FROM drivers WHERE email = $1", email)
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

func (repository *DriverRepository) GetAvailableDrivers() ([]models.Driver, error) {
	dateNow := time.Now().UTC().Format("2006-01-02")
	log.Println(dateNow)

	rows, err := repository.DataBase.Query(`
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
