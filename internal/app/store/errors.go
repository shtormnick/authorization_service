package store

// ErrRecordNotFound ...
type ErrRecordNotFound struct{}

func (e ErrRecordNotFound) Error() string {
	return "Record not found"
}
