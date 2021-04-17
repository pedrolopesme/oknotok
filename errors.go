package oknotok

import "errors"

var (
	// ErrCircuitNotOk is always returned when OkNotOk detects a dangerous amount of errors
	ErrCircuitNotOk = errors.New("circuit breaker is in not ok period, so calls are not going to happening")
	// ErrTooManyCalls is returned when OkNotOk is in HalfOk state and the calls count has reached MaxHalfOkRequests
	ErrTooManyCalls = errors.New("circuit break wont allow new requests at this time due its HalkOk state")
)
