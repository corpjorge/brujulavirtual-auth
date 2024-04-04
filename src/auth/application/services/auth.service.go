package services

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
)

type Impl struct {
	repo ports.Repository
}

func AuthService(repo ports.Repository) ports.Service {
	return &Impl{repo}
}

func (i *Impl) Create(auth models.Auth) (models.Auth, error) {
	_, err := i.repo.Create(auth)
	if err != nil {
		return models.Auth{}, err
	}
	return auth, nil
}
