package services

import (
	"brujulavirtual-auth/src/auth/domain/models"
	ports2 "brujulavirtual-auth/src/auth/domain/ports"
)

type Impl struct {
	repo ports2.Repository
}

func AuthService(repo ports2.Repository) ports2.Service {
	return &Impl{repo}
}

func (i *Impl) Create(auth models.Auth) (models.Auth, error) {
	_, err := i.repo.Create(auth)
	if err != nil {
		return models.Auth{}, err
	}
	return auth, nil
}
