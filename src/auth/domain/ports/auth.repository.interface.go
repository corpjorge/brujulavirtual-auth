package ports

import (
	"brujulavirtual-auth/src/auth/domain/models"
)

type Repository interface {
	Create(auth models.Auth) (models.Auth, error)
}
