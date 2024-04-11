package ports

import (
	"brujulavirtual-auth/src/auth/domain/models"
)

type Service interface {
	Validate(auth models.Auth) (models.Auth, error)
}
