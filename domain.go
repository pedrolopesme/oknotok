package oknotok

import (
	"fmt"
	"time"
)

// represents the current state of the circuit breaker
type CircuitState int

// default settings
const (
	defaultInterval              = time.Duration(0) * time.Second
	defaultTimeout               = time.Duration(60) * time.Second
	defaultMaxContinuousFailures = 5
)

// available states
const (
	StateOk CircuitState = iota
	StateHalfOk
	StateNotOk
)

// cast state to string
func (s CircuitState) String() string {
	switch s {
	case StateNotOk:
		return "not ok - circuit closed"
	case StateHalfOk:
		return "half ok - circuit half-open"
	case StateOk:
		return "ok - circuit open"
	default:
		return fmt.Sprintf("unknwon state: %d", s)
	}
}

// keeps the current status of the requests, their failures and successes
// TODO clear the internal status on the change of the state and/or at the
// close state intervals
type Stats struct {
	Requests            uint64
	Failures            uint64
	Successes           uint64
	ContinuousSuccesses uint64
	ContinuousFailures  uint64
}

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

// Core domain type, implementing the most important features of a
// CircuitBreaker. It helps to protect the environment
// from sending requests that probably are going to fail.
// Source of inspiration: https://martinfowler.com/bliki/CircuitBreaker.html
type OkNotOk struct {
	name         string
	maxRequests  uint64
	interval     time.Duration
	timeout      time.Duration
	readyToTrip  func(stats Stats) bool
	stateChanged func(name string, from, to CircuitState)
	isSuccessful func(err error) bool
}

// exposes name
func (ok *OkNotOk) Name() string {
	return ok.Name()
}
