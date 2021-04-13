package oknotok

import "time"

// general settings applyied to every CircuitBreaker
type Settings struct {
	Name              string                                   // identify the ciruit breaker
	MaxHalfOkRequests uint64                                   // maximum of requests allowed in half-open state
	Interval          time.Duration                            // interval to clear internal State during NotOk state
	Timeout           time.Duration                            // period of Ok state after the circuit breaker becomes halfOk
	Healed            func(stats Stats) bool                   // given a Stats, this func will check if the circuit break should change to Ok state
	StateChanged      func(name string, from, to CircuitState) // callback called whenever circuit state changes
	CountError        func(err error) bool                     // check if an error should be counted to the internal Stats counter
}
