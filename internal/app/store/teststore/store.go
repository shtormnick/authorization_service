package teststore

import (
	"github.com/shitikovkirill/auth-service/internal/app/store"
)

// Store ...
type Store struct {
}

// New ...
func New() *Store {
	return &Store{}
}

// Close db connection
func (s *Store) Close() error {
	return nil
}

// BeginTransaction ...
func (s *Store) BeginTransaction() (*store.SQLOptions, error) {
	return &store.SQLOptions{}, nil
}

// CommitTransaction ...
func (s *Store) CommitTransaction(opt *store.SQLOptions) error {
	return nil
}

// RollbackTransaction ...
func (s *Store) RollbackTransaction(opt *store.SQLOptions) error {
	return nil
}
