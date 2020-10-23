package handler

import (
	"net/http"
)

type responseInterface interface {
	JSON(w http.ResponseWriter, data interface{}, statusCode int)
	Error(w http.ResponseWriter, err error, statusCode int)
}
