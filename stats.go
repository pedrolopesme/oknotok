package go_oknotok

// keeps the current status of the requests, their failures and successes
// TODO clear the internal status on the change of the state and/or at the
// close state intervals
type Stats struct {
	calls               uint64
	failures            uint64
	successes           uint64
	continuousSuccesses uint64
	continuousFailures  uint64
}

// clears all internal stats
// TODO add tests
func (s *Stats) reset() {
	s.calls = 0
	s.failures = 0
	s.successes = 0
	s.continuousFailures = 0
	s.continuousSuccesses = 0
}

// increments a call
// TODO add tests
func (s *Stats) onCall() {
	s.calls++
}

// increments a success
// TODO add tests
func (s *Stats) onSuccess() {
	s.successes++
	s.continuousSuccesses++
	s.continuousFailures = 0
}

// increments a failure
// TODO add tests
func (s *Stats) onFailure() {
	s.failures++
	s.continuousFailures++
	s.continuousSuccesses = 0
}
