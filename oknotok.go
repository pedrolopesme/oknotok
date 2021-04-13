package oknotok

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
	return stats.ContinuousFailures > defaultMaxContinuousFailures
}

func defaultIsSuccessful(err error) bool {
	return err == nil
}

// trigger a request if the current state of
// OkNotOk allows it. In case of OkNotOk rejection,
// an error will be returned.
func (ok *OkNotOk) Call(req func() (interface{}, error)) (interface{}, error) {
	result, err := req()
	return result, err
}
