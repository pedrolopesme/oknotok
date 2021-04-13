package oknotok

import "time"

// default settings
const (
	defaultInterval              = time.Duration(0) * time.Second
	defaultTimeout               = time.Duration(60) * time.Second
	defaultMaxContinuousFailures = 5
)

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

// returns a new OkNotOk instance properly configured
func NewOkNotOk(settings Settings) *OkNotOk {
	oknok := OkNotOk{}
	oknok.name = settings.Name
	oknok.stateChanged = settings.StateChanged

	if settings.Interval > 0 {
		oknok.interval = settings.Interval
	} else {
		oknok.interval = defaultInterval
	}

	if settings.MaxRequests > 0 {
		oknok.maxRequests = settings.MaxRequests
	} else {
		oknok.maxRequests = 1
	}

	if settings.Timeout > 0 {
		oknok.timeout = settings.Timeout
	} else {
		oknok.timeout = defaultTimeout
	}

	if settings.IsSuccessful != nil {
		oknok.isSuccessful = settings.IsSuccessful
	} else {
		oknok.isSuccessful = defaultIsSuccessful
	}

	if settings.ReadyToTrip != nil {
		oknok.readyToTrip = settings.ReadyToTrip
	} else {
		oknok.readyToTrip = defaultReadyToTrip
	}

	return &oknok
}

func defaultReadyToTrip(stats Stats) bool {
	return stats.continuousFailures > defaultMaxContinuousFailures
}

func defaultIsSuccessful(err error) bool {
	return err == nil
}

// trigger a request if the current state of
// OkNotOk allows it. In case of OkNotOk rejection,
// an error will be returned.
// TODO implement PreTrip and Done
func (ok *OkNotOk) Call(req func() (interface{}, error)) (interface{}, error) {
	result, err := req()
	return result, err
}
