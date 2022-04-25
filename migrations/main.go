package migrations

import (
	"database/sql"
	"log"
	"root/config"
)

func Run() {
	migrate(config.DB, Drivers)
	migrate(config.DB, TripsDrop)
	migrate(config.DB, TripsCreation)
	migrate(config.DB, TripsMigration)
}
func migrate(dbDriver *sql.DB, query string) {
	statement, err := dbDriver.Prepare(query)
	if err == nil {
		_, creationError := statement.Exec()
		if creationError == nil {
			log.Println("BD updated OK")
		} else {
			log.Println(creationError.Error())
		}
	} else {
		log.Println(err.Error())
	}
}
