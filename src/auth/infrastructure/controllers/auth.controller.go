package controllers

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
	"brujulavirtual-auth/src/commons"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	service ports.Service
}

func AuthController(service ports.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Validate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.ValidatePost(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (c *Controller) ValidatePost(w http.ResponseWriter, r *http.Request) {

	var auth models.Auth
	err := json.NewDecoder(r.Body).Decode(&auth)
	log.Default().Println(auth)
	if err != nil {
		commons.ErrorResponse(w, "Error processing data", http.StatusBadRequest)
		return
	}

	if !auth.IsValid() {
		commons.ErrorResponse(w, "Incomplete or invalid authentication data", http.StatusBadRequest)
		return
	}

	createdAuth, err := c.service.Validate(auth)
	if err != nil {
		commons.ErrorResponse(w, "Error creating authenticator entity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdAuth)
	if err != nil {
		log.Default().Printf("Error encoding the response: %v", err)
	}

}
