package fastcache

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func NewRealClock() Clock {
	return RealClock{}
}

func (rc RealClock) Now() time.Time {
	t := time.Now()
	return t
}
