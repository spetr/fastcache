package fastcache

import "errors"

var (
	// ErrKeyNotFound - error retirned when requested key was not found i cache
	ErrKeyNotFound = errors.New("key not found")
)
