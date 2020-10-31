package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() http.Handler {
	httpHandler := GetInjectedHandler()
	r := mux.NewRouter()
	r.HandleFunc("/write", httpHandler.WriteHandler).Methods("POST")
	return r
}
