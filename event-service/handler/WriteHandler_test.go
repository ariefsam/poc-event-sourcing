package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ariefsam/poc-event-sourcing/event-service/handler"
)

func TestWriteHandler(t *testing.T) {
	log.Println("time", time.Now().Unix())
	time.Sleep(100 * time.Millisecond)
	log.Println(time.Now().UnixNano())
	event := map[string]interface{}{
		"EventName": "UserCreated",
		"Payload": map[string]interface{}{
			"ID":    "j10",
			"Name":  "Arief Hidayatulloh",
			"Email": "arief.hidayatulloh@gmail.com",
			"Role":  "User",
			"CreatedBy": map[string]interface{}{
				"ID":    "j1",
				"Name":  "Nam Indra",
				"Email": "nam@jojonomic.com",
				"Role":  "Admin",
			},
			"CreatedAtTimestamp": 1603229314,
		},
	}

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
