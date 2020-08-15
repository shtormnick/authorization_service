package realservice

import (
	"github.com/shitikovkirill/auth-service/internal/app/store"
)

// Service ...
type Service struct {
	store store.Store
}

// New ...
func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}
