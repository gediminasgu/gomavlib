//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"errors"
)

// States of the mission state machine.
// Note that these states are independent of whether the mission is in a mode that can execute mission items or not (is suspended).
// They may not all be relevant on all vehicles.
type MISSION_STATE uint32

const (
	// The mission status reporting is not supported.
	MISSION_STATE_UNKNOWN MISSION_STATE = 0
	// No mission on the vehicle.
	MISSION_STATE_NO_MISSION MISSION_STATE = 1
	// Mission has not started. This is the case after a mission has uploaded but not yet started executing.
	MISSION_STATE_NOT_STARTED MISSION_STATE = 2
	// Mission is active, and will execute mission items when in auto mode.
	MISSION_STATE_ACTIVE MISSION_STATE = 3
	// Mission is paused when in auto mode.
	MISSION_STATE_PAUSED MISSION_STATE = 4
	// Mission has executed all mission items.
	MISSION_STATE_COMPLETE MISSION_STATE = 5
)

var labels_MISSION_STATE = map[MISSION_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MISSION_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_MISSION_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MISSION_STATE = map[string]MISSION_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MISSION_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MISSION_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MISSION_STATE) String() string {
	if l, ok := labels_MISSION_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
