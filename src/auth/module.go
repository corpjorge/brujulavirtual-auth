package auth

import (
	"brujulavirtual-auth/src/auth/application/services"
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"brujulavirtual-auth/src/auth/infrastructure/repositories"
	"brujulavirtual-auth/src/auth/infrastructure/routes"
	"net/http"
)

func Module(mux *http.ServeMux) {
	authRepository := repositories.Auth()
	authService := services.Auth(authRepository)
	authController := controllers.Auth(authService)

	routes.Router(*authController, mux)
}
