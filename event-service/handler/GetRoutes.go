package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/write", WriteHandler).Methods("POST", "GET")
	return r
}
