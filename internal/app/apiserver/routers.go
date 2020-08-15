package apiserver

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	s.router.HandleFunc("/api/status", s.handleStatus()).Methods("GET")

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	s.router.Handle("/docs", sh)
	s.router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

}
