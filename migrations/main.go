package migrations

import (
	"database/sql"
	"log"
	"root/config"
)

func Run() {
	// Migrate drivers
	migrate(config.DB, Drivers)
	// Other migrations can be added here.
}
func migrate(dbDriver *sql.DB, query string) {
	statement, err := dbDriver.Prepare(query)
	if err == nil {
		_, creationError := statement.Exec()
		if creationError == nil {
			log.Println("Table created successfully")
		} else {
			log.Println(creationError.Error())
		}
	} else {
		log.Println(err.Error())
	}
}
