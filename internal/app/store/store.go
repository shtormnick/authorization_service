package store

// Store ...
type Store interface {
	Close() error
	BeginTransaction() (*SQLOptions, error)
	CommitTransaction(*SQLOptions) error
	RollbackTransaction(*SQLOptions) error
}
