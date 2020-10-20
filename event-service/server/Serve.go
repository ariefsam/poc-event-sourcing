package server

import (
	"log"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/handlers"
)

func Serve(port string, r http.Handler) {
	gzip := gziphandler.GzipHandler(r)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        handlers.CORS()(gzip),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
