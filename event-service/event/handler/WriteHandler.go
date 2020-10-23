package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HttpHandler) WriteHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	json.NewDecoder(r.Body).Decode(&payload)
	id, err := h.EventService.Write(payload)
	if err != nil {
		h.response.Error(w, err, http.StatusBadGateway)
	}
	h.response.JSON(w, id, http.StatusOK)
}
