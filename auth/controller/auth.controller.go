package controller

import (
	"brujulavirtual-auth/auth/models"
	"brujulavirtual-auth/auth/service"
	"encoding/json"
	"net/http"
)

type Controller struct {
	service service.Service
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.CreatePost(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (c *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {

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

	createdAuth, err := c.service.Create(auth)
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

func AuthController(service service.Service) *Controller {
	return &Controller{service: service}
}
