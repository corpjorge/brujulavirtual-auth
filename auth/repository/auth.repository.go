package repository

import (
	"brujulavirtual-auth/auth/models"
)

type Impl struct {
}

func (r Impl) Create(auth models.Auth) (models.Auth, error) {
	return auth, nil
}

func AuthRepository() Repository {
	return &Impl{}
}
