package repository

import "brujulavirtual-auth/auth/models"

type Repository interface {
	Create(auth models.Auth) (models.Auth, error)
}
