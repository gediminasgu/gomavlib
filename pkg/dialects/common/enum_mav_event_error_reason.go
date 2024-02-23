//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Reason for an event error response.
type MAV_EVENT_ERROR_REASON uint32

const (
	// The requested event is not available (anymore).
	MAV_EVENT_ERROR_REASON_UNAVAILABLE MAV_EVENT_ERROR_REASON = 0
)

var labels_MAV_EVENT_ERROR_REASON = map[MAV_EVENT_ERROR_REASON]string{
	MAV_EVENT_ERROR_REASON_UNAVAILABLE: "MAV_EVENT_ERROR_REASON_UNAVAILABLE",
}

var values_MAV_EVENT_ERROR_REASON = map[string]MAV_EVENT_ERROR_REASON{
	"MAV_EVENT_ERROR_REASON_UNAVAILABLE": MAV_EVENT_ERROR_REASON_UNAVAILABLE,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_EVENT_ERROR_REASON) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_EVENT_ERROR_REASON[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_EVENT_ERROR_REASON) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_EVENT_ERROR_REASON[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = MAV_EVENT_ERROR_REASON(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_EVENT_ERROR_REASON) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
