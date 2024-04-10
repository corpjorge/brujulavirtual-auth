package ports

import (
	"brujulavirtual-auth/src/auth/domain/models"
)

type Repository interface {
	Validate(auth models.Auth) (models.Auth, error)
}
