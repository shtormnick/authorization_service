package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/shitikovkirill/auth-service/internal/app/store"
)

// Store ...
type Store struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Close db connection
func (s *Store) Close() error {
	return s.db.Close()
}

// BeginTransaction function
func (s *Store) BeginTransaction() (*store.SQLOptions, error) {
	tx, err := s.db.Begin()
	return &store.SQLOptions{
		Transaction: tx,
	}, err
}

// CommitTransaction ...
func (s *Store) CommitTransaction(opt *store.SQLOptions) error {
	if opt.Transaction == nil {
		return errors.New("You try Commit not oppen transaction")
	}
	err := opt.Transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

// RollbackTransaction ...
func (s *Store) RollbackTransaction(opt *store.SQLOptions) error {
	if opt.Transaction == nil {
		return errors.New("You try Rollback not oppen transaction")
	}
	err := opt.Transaction.Rollback()
	if err != nil {
		return err
	}
	return nil
}
