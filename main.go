package main

import (
	"brujulavirtual-auth/auth/config"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	config.AuthModule(mux)

	fmt.Println("The server is listening on the port http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
	}
}
