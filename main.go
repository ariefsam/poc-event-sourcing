package main

import (
	"github.com/ariefsam/poc-event-sourcing/event-service/handler"
	"github.com/ariefsam/poc-event-sourcing/event-service/server"
)

func main() {
	routes := handler.GetRoutes()
	server.Serve("8080", routes)
}
