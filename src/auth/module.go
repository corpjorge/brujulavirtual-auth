package auth

import (
	"brujulavirtual-auth/src/auth/application/services"
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"brujulavirtual-auth/src/auth/infrastructure/repositories"
	"brujulavirtual-auth/src/auth/infrastructure/routes"
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
