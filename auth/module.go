package auth

import (
	"brujulavirtual-auth/auth/application/services"
	"brujulavirtual-auth/auth/infrastructure/controllers"
	"brujulavirtual-auth/auth/infrastructure/repositories"
	"brujulavirtual-auth/auth/infrastructure/routes"
	"net/http"
)

func Setup() *controllers.Controller {
	authRepository := repositories.AuthRepository()
	authService := services.AuthService(authRepository)
	authController := controllers.AuthController(authService)

	return authController

}

func AuthModule(mux *http.ServeMux) {
	authController := Setup()
	routes.Router(*authController, mux)
}
