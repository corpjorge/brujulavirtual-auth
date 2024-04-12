package auth

import (
	"brujulavirtual-auth/src/auth/domain/ports"
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"brujulavirtual-auth/src/auth/infrastructure/repositories"
	"brujulavirtual-auth/src/auth/infrastructure/routes"
	"net/http"
)

func Module(mux *http.ServeMux) {
	authRepository := repositories.Auth()
	authService := ports.Service(authRepository)
	authController := controllers.Auth(authService)

	routes.Router(*authController, mux)
}
