package controllers

type AuthService interface {
	CheckPassword(providedPassword string, anotherPass string) error
	HashPassword(password string) (string, error)
}
