package controller

import (
	"context"
	"dttg/internal/service"
	"dttg/pkg/api"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SpellController struct {
	service service.Service
}

func NewSpellController(s service.Service) *SpellController {
	return &SpellController{
		service: s,
	}
}

func (ac *SpellController) GetAll(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	res, err := ac.service.GetAllServ(ctx)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *SpellController) Get(w http.ResponseWriter, r *http.Request) {
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

func (ac *SpellController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/spell", ac.GetAll).Methods("GET")
	router.HandleFunc("/spell/{name}", ac.Get).Methods("GET")
}
