package main

import (
	"jojoevent/error_logger"
	"jojoevent/event/repository"
	"log"
	"os"

	"github.com/keimoon/gore"

	"github.com/ariefsam/poc-event-sourcing/event-service/event/handler"

	"github.com/ariefsam/poc-event-sourcing/event-service/server"
)

func main() {
	routes := GetRoutes()
	server.Serve("8001", routes)
}

// Get handler that injected with implementation service
func GetInjectedHandler() (h handler.HttpHandler) {
	var eventService handler.EventService
	db := os.Getenv("DB_DRIVER")

	if db == "redis" {
		conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
		if err != nil {
			log.Fatal(err)
		}

		eventService = &repository.RedisEvent{
			Conn: conn,
		}
	} else {
		eventService = &repository.MySQLEvent{}
	}

	h.EventService = eventService
	h.ErrorLog = &error_logger.ErrorPrinter{}
	return
}
