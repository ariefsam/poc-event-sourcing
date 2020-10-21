package main

import (
	"net/http"

	"github.com/ariefsam/poc-event-sourcing/event-service/handler"
	"github.com/ariefsam/poc-event-sourcing/event-service/server"
	"github.com/gorilla/mux"
)

func main() {
	routes := GetRoutes()
	server.Serve("8080", routes)
}

func GetRoutes() http.Handler {
	httpHandler := GetInjectedHandler()
	r := mux.NewRouter()
	r.HandleFunc("/write", httpHandler.WriteHandler).Methods("POST")
	return r
}

// Get handler that injected with implementation service
func GetInjectedHandler() (h handler.HttpHandler) {

	return
}
