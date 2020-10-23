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
		},
	}
	readEvents := []map[string]interface{}{
		map[string]interface{}{
			"EventName": "UserCreated",
			"Payload": map[string]interface{}{
				"ID":    "j10",
				"Name":  "Arief Hidayatulloh",
				"Phone": "0852163772",
				"Email": "arief.hidayatulloh@gmail.com",
				"Role":  "User",
				"CreatedBy": map[string]interface{}{
					"ID":    "j1",
					"Name":  "Nam Indra",
					"Email": "nam@jojonomic.com",
					"Role":  "Admin",
				},
				"CreatedAtTimestamp": 1603235393,
			},
			"SystemUnixNano": 1603235393816625000,
		},
		map[string]interface{}{
			"EventName": "UserUpdated",
			"Payload": map[string]interface{}{
				"ID":    "j10",
				"Name":  "Arief Hidayatulloh S.Komp",
				"Phone": "085212345678",
				"Email": "arief.hidayatulloh@gmail.com",
				"Role":  "User",
				"UpdatedBy": map[string]interface{}{
					"ID":    "j1",
					"Name":  "Nam Indra",
					"Email": "nam@jojonomic.com",
					"Role":  "Admin",
				},
				"CreatedAtTimestamp": 1603235414,
			},
			"SystemUnixNano": 1603235414254786000,
		},
	}

	requestJson, _ := json.Marshal(event)

	r := httptest.NewRequest("GET", "http://xx/x", bytes.NewReader(requestJson))
	w := httptest.NewRecorder()

	mockEventService := GetEventServiceMock(readEvents)
	h := handler.HttpHandler{
		EventService: &mockEventService,
	}
	h.WriteHandler(w, r)

	mockEventService.AssertCalled(t, "Write", event)
}
