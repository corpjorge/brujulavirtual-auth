package auth

import (
	"brujulavirtual-auth/src/auth/application/services"
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"brujulavirtual-auth/src/auth/infrastructure/repositories"
	"brujulavirtual-auth/src/auth/infrastructure/routes"
	"net/http"
)

func Module(mux *http.ServeMux) {
	authRepository := repositories.AuthRepository()
	authService := services.AuthService(authRepository)
	authController := controllers.AuthController(authService)

	routes.Router(*authController, mux)
}
