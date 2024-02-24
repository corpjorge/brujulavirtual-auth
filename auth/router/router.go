package router

import (
	"brujulavirtual-auth/auth/controller"
	"net/http"
)

func Router(auth controller.Controller, mux *http.ServeMux) {
	mux.HandleFunc("/auth", auth.Create)
}
