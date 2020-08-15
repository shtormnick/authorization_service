package db

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/shitikovkirill/auth-service/internal/app/store/sqlstore"
)

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// GetStore ...
func GetStore(dbURL string) (*sqlstore.Store, error) {
	db, err := newDB(dbURL)
	if err != nil {
		return nil, err
	}

	return sqlstore.New(db), nil
}
