package store

import (
	"database/sql"
)

// SQLOptions Parameter that contain transaction and another options
type SQLOptions struct {
	Transaction *sql.Tx
	ForUpdate   bool
}
