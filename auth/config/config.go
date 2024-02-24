package config

import (
	"brujulavirtual-auth/auth/controller"
	"brujulavirtual-auth/auth/repository"
	"brujulavirtual-auth/auth/router"
	"brujulavirtual-auth/auth/service"
	"net/http"
)

func Setup() *controller.Controller {
	authRepository := repository.AuthRepository()
	authService := service.AuthService(authRepository)
	authController := controller.AuthController(authService)

	return authController

}

func AuthModule(mux *http.ServeMux) {
	authController := Setup()
	router.Router(*authController, mux)
}
