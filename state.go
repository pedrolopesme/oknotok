package oknotok

import "fmt"

// represents the current state of the circuit breaker
type CircuitState int

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
