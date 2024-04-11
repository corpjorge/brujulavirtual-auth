package controllers

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
	"encoding/json"
	"net/http"
)

type Controller struct {
	service ports.Service
}

type ErrorResponse struct {
	Error string `json:"error"`
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
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Error: "Error processing data",
		}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	if !auth.IsValid() {
		http.Error(w, "Incomplete or invalid authentication data", http.StatusBadRequest)
		return
	}

	createdAuth, err := c.service.Validate(auth)
	if err != nil {
		http.Error(w, "Error creating authenticator entity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdAuth)
	if err != nil {
		return
	}

}
