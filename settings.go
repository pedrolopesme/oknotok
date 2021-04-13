package oknotok

import "time"

// general settings applyied to every CircuitBreaker
type Settings struct {
	Name         string
	MaxRequests  uint64
	Interval     time.Duration
	Timeout      time.Duration
	ReadyToTrip  func(stats Stats) bool
	StateChanged func(name string, from, to CircuitState)
	IsSuccessful func(err error) bool
}
