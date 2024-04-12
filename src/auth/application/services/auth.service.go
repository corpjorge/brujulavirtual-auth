package services

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Service struct {
	repository ports.Repository
}

func (service *Service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (service *Service) Validate(auth models.Auth) (models.Auth, error) {
	storedAuth, err := service.repository.Validate(auth)
	log.Default().Println("storedAuth ", storedAuth)
	if err != nil {
		return models.Auth{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedAuth.Password), []byte(auth.Password))
	if err != nil {
		return models.Auth{}, err
	}

	return auth, nil
}
