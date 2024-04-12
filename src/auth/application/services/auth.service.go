package services

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
)

type Service struct {
	repository ports.Repository
}

func (service *Service) Validate(auth models.Auth) (models.Auth, error) {
	_, err := service.repository.Validate(auth)
	if err != nil {
		return models.Auth{}, err
	}
	return auth, nil
}
