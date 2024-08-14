package infrastructure

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)



type PasswordService interface {
	HashPassword(string) (string, error)
	ComparePassword(string, string) error
}


type passwordService struct {}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (ps *passwordService) HashPassword(password string) (string, error) {
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Failed to hash password")
	}
	return string(hashedPassword), nil
}

func (ps *passwordService) ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("Invalid password")
	}
	return nil
}


