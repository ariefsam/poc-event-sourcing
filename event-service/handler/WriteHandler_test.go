package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ariefsam/poc-event-sourcing/event-service/handler"
)

func TestWriteHandler(t *testing.T) {
	event := map[string]interface{}{
		"EventName": "UserCreated",
		"Payload": map[string]interface{}{
			"ID":    "j10",
			"Name":  "Arief Hidayatulloh",
			"Email": "arief.hidayatulloh@gmail.com",
			"Role":  "User",
			"Creator": map[string]interface{}{
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
	handler.WriteHandler(w, r)

	var mockEventService EventServiceMock
	mockEventService.AssertCalled(t, "Write", event)
}
