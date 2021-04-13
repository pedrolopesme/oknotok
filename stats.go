package oknotok

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
