package main

import (
	"log"
	"net/http"

	"github.com/ariefsam/poc-event-sourcing/event-service/event/error_logger"
	"github.com/keimoon/gore"

	"github.com/ariefsam/poc-event-sourcing/event-service/event/handler"
	"github.com/ariefsam/poc-event-sourcing/event-service/event/repository"
	"github.com/ariefsam/poc-event-sourcing/event-service/server"
	"github.com/gorilla/mux"
)

func main() {
	routes := GetRoutes()
	server.Serve("8001", routes)
}

func GetRoutes() http.Handler {
	httpHandler := GetInjectedHandler()
	r := mux.NewRouter()
	r.HandleFunc("/write", httpHandler.WriteHandler).Methods("POST")
	return r
}

// Get handler that injected with implementation service
func GetInjectedHandler() (h handler.HttpHandler) {
	conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
	if err != nil {
		log.Fatal(err)
	}

	eventService := &repository.LedisEvent{
		Conn: conn,
	}

	h.EventService = eventService
	h.ErrorLog = &error_logger.ErrorPrinter{}
	return
}
