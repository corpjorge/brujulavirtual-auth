package ports

import (
	"brujulavirtual-auth/auth/domain/models"
)

type Service interface {
	Create(auth models.Auth) (models.Auth, error)
}
