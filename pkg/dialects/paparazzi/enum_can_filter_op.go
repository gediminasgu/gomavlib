//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package paparazzi

import (
	"errors"
)

type CAN_FILTER_OP uint32

const (
	CAN_FILTER_REPLACE CAN_FILTER_OP = 0
	CAN_FILTER_ADD     CAN_FILTER_OP = 1
	CAN_FILTER_REMOVE  CAN_FILTER_OP = 2
)

var labels_CAN_FILTER_OP = map[CAN_FILTER_OP]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAN_FILTER_OP) MarshalText() ([]byte, error) {
	if l, ok := labels_CAN_FILTER_OP[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAN_FILTER_OP = map[string]CAN_FILTER_OP{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAN_FILTER_OP) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAN_FILTER_OP[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAN_FILTER_OP) String() string {
	if l, ok := labels_CAN_FILTER_OP[e]; ok {
		return l
	}
	return "invalid value"
}
