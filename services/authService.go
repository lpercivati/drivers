package services

import "golang.org/x/crypto/bcrypt"

type AuthService struct {
}

func (service *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (service *AuthService) CheckPassword(providedPassword string, anotherPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(anotherPass), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
