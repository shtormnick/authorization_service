package dto

import (
	"net/http"
)

// ResponseWriter ...
type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

// WriteHeader ...
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// StatusResponse - return status
type StatusResponse struct {
	// Return "ok" if api work
	Status string `json:"status"`
}
