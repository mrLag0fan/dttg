package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	HandleRequests(router *mux.Router)
}

func ErrorResponse(w *http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err.Error())
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}
