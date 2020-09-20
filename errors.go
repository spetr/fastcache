package fastcache

import "errors"

var (
	// ErrKeyNotFoundError - error retirned when requested key was not found i cache
	ErrKeyNotFoundError = errors.New("key not found")
)
