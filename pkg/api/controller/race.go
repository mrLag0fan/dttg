package controller

import (
	"context"
	"dttg/internal/service"
	"dttg/pkg/api"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type RaceController struct {
	service service.Service
}

func NewRaceController(s service.Service) *RaceController {
	return &RaceController{
		service: s,
	}
}

func (ac *RaceController) GetAll(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	res, err := ac.service.GetAllServ(ctx)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *RaceController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["name"]
	res, err := ac.service.GetByNameServ(ctx, id)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *RaceController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/race", ac.GetAll).Methods("GET")
	router.HandleFunc("/race/{name}", ac.Get).Methods("GET")
}
