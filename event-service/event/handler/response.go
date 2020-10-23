package handler

import (
	"encoding/json"
	"net/http"
)

type response struct{}

func (r *response) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	view, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(view)
}

func (r *response) Error(w http.ResponseWriter, err error, statusCode int) {
	response := map[string]interface{}{}
	if err != nil {
		response["Error"] = err.Error()
		r.JSON(w, response, statusCode)
		return
	}
}
