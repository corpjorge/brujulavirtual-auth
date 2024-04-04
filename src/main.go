package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	App(mux)

	fmt.Println("The server is listening on the port http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
