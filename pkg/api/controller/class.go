package controller

import (
	"context"
	"dttg/internal/service"
	"dttg/pkg/api"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ClassController struct {
	service service.Service
}

func NewClassController(s service.Service) *ClassController {
	return &ClassController{
		service: s,
	}
}

func (ac *ClassController) GetAll(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	res, err := ac.service.GetAllClasses(ctx)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *ClassController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["name"]
	res, err := ac.service.GetByClassName(ctx, id)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *ClassController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/class", ac.GetAll).Methods("GET")
	router.HandleFunc("/class/{name}", ac.Get).Methods("GET")
}
