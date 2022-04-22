package migrations

const Drivers = `
CREATE TABLE IF NOT EXISTS drivers (
  Id INTEGER PRIMARY KEY AUTOINCREMENT,
  Fullname VARCHAR(64) NOT NULL,
  Email VARCHAR(64) NOT NULL,
  PasswordHash VARCHAR(256) NOT NULL,
  IsAdmin bit NOT NULL,
  DateCreation TIMESTAMP NOT NULL
)
`
