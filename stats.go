package oknotok

// keeps the current status of the requests, their failures and successes
// TODO clear the internal status on the change of the state and/or at the
// close state intervals
type Stats struct {
	requests            uint64
	failures            uint64
	successes           uint64
	continuousSuccesses uint64
	continuousFailures  uint64
}

// clear all internal stats
// TODO add tests
func (s *Stats) reset() {
	s.requests = 0
	s.failures = 0
	s.successes = 0
	s.continuousFailures = 0
	s.continuousSuccesses = 0
}
