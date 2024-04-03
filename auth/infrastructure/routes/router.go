package routes

import (
	"brujulavirtual-auth/auth/infrastructure/controllers"
	"net/http"
)

func Router(auth controllers.Controller, mux *http.ServeMux) {
	mux.HandleFunc("/auth", auth.Create)
}
