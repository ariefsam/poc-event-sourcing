package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ariefsam/poc-event-sourcing/event-service/event/handler"
)

func TestWriteHandler(t *testing.T) {
	log.Println("time", time.Now().Unix())
	time.Sleep(100 * time.Millisecond)
	log.Println(time.Now().UnixNano())
	event := GetSampleEvent()
	requestJson, _ := json.Marshal(event)

	r := httptest.NewRequest("GET", "http://xx/x", bytes.NewReader(requestJson))
	w := httptest.NewRecorder()

	mockEventService := GetEventServiceMock()
	h := handler.HttpHandler{
		EventService: &mockEventService,
	}
	h.WriteHandler(w, r)

	mockEventService.AssertCalled(t, "Write", event)
}
