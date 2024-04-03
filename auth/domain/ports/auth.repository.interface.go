package ports

import (
	"brujulavirtual-auth/auth/domain/models"
)

type Repository interface {
	Create(auth models.Auth) (models.Auth, error)
}
