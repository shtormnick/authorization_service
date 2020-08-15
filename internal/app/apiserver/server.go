package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/shitikovkirill/auth-service/internal/app/service"
)

type server struct {
	router  *mux.Router
	service service.Service
}

func newServer(service service.Service) *server {
	s := &server{
		router:  mux.NewRouter(),
		service: service,
	}

	s.configureRouter()

	return s
}
