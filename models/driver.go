package models

import "time"

type Driver struct {
	Id           int
	Fullname     string
	Email        string
	PasswordHash string
	IsAdmin      bool
	DateCreation time.Time
}
