package auth

import (
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"brujulavirtual-auth/src/auth/infrastructure/repositories"
	"brujulavirtual-auth/src/auth/infrastructure/routes"
	"net/http"
)

func Module(mux *http.ServeMux) {
	authRepository := repositories.AuthRepository()
	//authService := ports.AuthService(authRepository)
	authController := controllers.Auth(authRepository)

	routes.Router(*authController, mux)
}
