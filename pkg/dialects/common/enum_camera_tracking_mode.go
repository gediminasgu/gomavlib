//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Camera tracking modes
type CAMERA_TRACKING_MODE uint32

const (
	// Not tracking
	CAMERA_TRACKING_MODE_NONE CAMERA_TRACKING_MODE = 0
	// Target is a point
	CAMERA_TRACKING_MODE_POINT CAMERA_TRACKING_MODE = 1
	// Target is a rectangle
	CAMERA_TRACKING_MODE_RECTANGLE CAMERA_TRACKING_MODE = 2
)

var labels_CAMERA_TRACKING_MODE = map[CAMERA_TRACKING_MODE]string{
	CAMERA_TRACKING_MODE_NONE:      "CAMERA_TRACKING_MODE_NONE",
	CAMERA_TRACKING_MODE_POINT:     "CAMERA_TRACKING_MODE_POINT",
	CAMERA_TRACKING_MODE_RECTANGLE: "CAMERA_TRACKING_MODE_RECTANGLE",
}

var values_CAMERA_TRACKING_MODE = map[string]CAMERA_TRACKING_MODE{
	"CAMERA_TRACKING_MODE_NONE":      CAMERA_TRACKING_MODE_NONE,
	"CAMERA_TRACKING_MODE_POINT":     CAMERA_TRACKING_MODE_POINT,
	"CAMERA_TRACKING_MODE_RECTANGLE": CAMERA_TRACKING_MODE_RECTANGLE,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_TRACKING_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_CAMERA_TRACKING_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_TRACKING_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_CAMERA_TRACKING_MODE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = CAMERA_TRACKING_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e CAMERA_TRACKING_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
