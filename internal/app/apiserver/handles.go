package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/shitikovkirill/auth-service/internal/app/apiserver/dto"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/logger"
)

// Test method for check that API run
func (s *server) handleStatus() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := &dto.StatusResponse{Status: "ok"}
		respondJSON(w, r, http.StatusOK, resp)
	})
}

func respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	logs := logger.Get()
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("Internal server error")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func respondError(w http.ResponseWriter, r *http.Request, err error) {
	logs := logger.Get()
	sentry.CaptureException(err)
	logs.Infof("RequestId: %s. Message: %v.", r.Context().Value(ctxKeyRequestID), err)
}

func respondJSONError(w http.ResponseWriter, r *http.Request, err error) {
	respondError(w, r, err)
	respondJSON(w, r, http.StatusInternalServerError, "Server Error")
}
