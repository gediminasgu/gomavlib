//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package matrixpilot

import (
	"errors"
)

// Winch status flags used in WINCH_STATUS
type MAV_WINCH_STATUS_FLAG uint32

const (
	// Winch is healthy
	MAV_WINCH_STATUS_HEALTHY MAV_WINCH_STATUS_FLAG = 1
	// Winch line is fully retracted
	MAV_WINCH_STATUS_FULLY_RETRACTED MAV_WINCH_STATUS_FLAG = 2
	// Winch motor is moving
	MAV_WINCH_STATUS_MOVING MAV_WINCH_STATUS_FLAG = 4
	// Winch clutch is engaged allowing motor to move freely.
	MAV_WINCH_STATUS_CLUTCH_ENGAGED MAV_WINCH_STATUS_FLAG = 8
	// Winch is locked by locking mechanism.
	MAV_WINCH_STATUS_LOCKED MAV_WINCH_STATUS_FLAG = 16
	// Winch is gravity dropping payload.
	MAV_WINCH_STATUS_DROPPING MAV_WINCH_STATUS_FLAG = 32
	// Winch is arresting payload descent.
	MAV_WINCH_STATUS_ARRESTING MAV_WINCH_STATUS_FLAG = 64
	// Winch is using torque measurements to sense the ground.
	MAV_WINCH_STATUS_GROUND_SENSE MAV_WINCH_STATUS_FLAG = 128
	// Winch is returning to the fully retracted position.
	MAV_WINCH_STATUS_RETRACTING MAV_WINCH_STATUS_FLAG = 256
	// Winch is redelivering the payload. This is a failover state if the line tension goes above a threshold during RETRACTING.
	MAV_WINCH_STATUS_REDELIVER MAV_WINCH_STATUS_FLAG = 512
	// Winch is abandoning the line and possibly payload. Winch unspools the entire calculated line length. This is a failover state from REDELIVER if the number of attempts exceeds a threshold.
	MAV_WINCH_STATUS_ABANDON_LINE MAV_WINCH_STATUS_FLAG = 1024
	// Winch is engaging the locking mechanism.
	MAV_WINCH_STATUS_LOCKING MAV_WINCH_STATUS_FLAG = 2048
	// Winch is spooling on line.
	MAV_WINCH_STATUS_LOAD_LINE MAV_WINCH_STATUS_FLAG = 4096
	// Winch is loading a payload.
	MAV_WINCH_STATUS_LOAD_PAYLOAD MAV_WINCH_STATUS_FLAG = 8192
)

var labels_MAV_WINCH_STATUS_FLAG = map[MAV_WINCH_STATUS_FLAG]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_WINCH_STATUS_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_WINCH_STATUS_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_WINCH_STATUS_FLAG = map[string]MAV_WINCH_STATUS_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_WINCH_STATUS_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_WINCH_STATUS_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_WINCH_STATUS_FLAG) String() string {
	if l, ok := labels_MAV_WINCH_STATUS_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}
