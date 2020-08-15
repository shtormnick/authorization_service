package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/shitikovkirill/auth-service/internal/app/apiserver/dto"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/logger"
	"github.com/shitikovkirill/auth-service/pkg/uuid"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyRequestID ctxKey = iota
)

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := uuid.GetUUID()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs := logger.Get()
		ctxLogger := logs.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		ctxLogger.Infof("started %s params: %v %s", r.Method, r.Form, r.RequestURI)

		start := time.Now()
		rw := &dto.ResponseWriter{ResponseWriter: w, Code: http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.Code >= 500:
			level = logrus.ErrorLevel
		case rw.Code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		ctxLogger.Logf(
			level,
			"completed with %d %s in %v",
			rw.Code,
			http.StatusText(rw.Code),
			time.Since(start),
		)
	})
}
