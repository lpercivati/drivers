package migrations

const TripsCreation = `
CREATE TABLE IF NOT EXISTS trips (
  DriverId INTEGER NOT NULL,
  DateStart TIMESTAMP NOT NULL,
  DateEnd TIMESTAMP NOT NULL
)
`

const TripsDrop = `
DROP TABLE Trips
`

const TripsMigration = `
INSERT INTO Trips values (1, '2022-04-24', '2022-10-20')
`
