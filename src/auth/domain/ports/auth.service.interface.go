package ports

import (
	"brujulavirtual-auth/src/auth/domain/models"
)

type AuthService interface {
	Validate(auth models.Auth) (models.Auth, error)
}
