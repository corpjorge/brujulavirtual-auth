package service

import (
	"brujulavirtual-auth/auth/models"
	"brujulavirtual-auth/auth/repository"
)

type Service interface {
	Create(auth models.Auth) (models.Auth, error)
}

type Impl struct {
	repo repository.Repository
}

func (i Impl) Create(auth models.Auth) (models.Auth, error) {
	_, err := i.repo.Create(auth)
	if err != nil {
		return models.Auth{}, err
	}
	return auth, nil
}

func AuthService(repo repository.Repository) Service {
	return &Impl{repo}
}
