package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HttpHandler) WriteHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	json.NewDecoder(r.Body).Decode(&payload)
	h.EventService.Write(payload)
}
