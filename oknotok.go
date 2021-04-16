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
	name              string
	maxHalfOkRequests uint64
	interval          time.Duration
	timeout           time.Duration
	healed            func(stats Stats) bool
	stateChanged      func(name string, from, to CircuitState)
	shouldCountError  func(err error) bool
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

	if settings.MaxHalfOkRequests > 0 {
		oknok.maxHalfOkRequests = settings.MaxHalfOkRequests
	} else {
		oknok.maxHalfOkRequests = 1
	}

	if settings.Timeout > 0 {
		oknok.timeout = settings.Timeout
	} else {
		oknok.timeout = defaultTimeout
	}

	if settings.ShoulCountError != nil {
		oknok.shouldCountError = settings.ShoulCountError
	} else {
		oknok.shouldCountError = defaultShouldCountError
	}

	if settings.Healed != nil {
		oknok.healed = settings.Healed
	} else {
		oknok.healed = defaultHealed
	}

	return &oknok
}

func defaultHealed(stats Stats) bool {
	return stats.continuousFailures > defaultMaxContinuousFailures
}

func defaultShouldCountError(err error) bool {
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
