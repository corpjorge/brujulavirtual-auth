package routes

import (
	"brujulavirtual-auth/src/auth/infrastructure/controllers"
	"net/http"
)

func Router(auth controllers.Controller, mux *http.ServeMux) {
	mux.HandleFunc("/auth", auth.Create)
	mux.HandleFunc("/authx", auth.Create)
}
