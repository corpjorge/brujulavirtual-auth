package main

import (
	"brujulavirtual-auth/src/auth"
	"net/http"
)

func App(mux *http.ServeMux) {
	auth.Module(mux)
}
