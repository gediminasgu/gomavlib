//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

type ACCELCAL_VEHICLE_POS uint32

const (
	ACCELCAL_VEHICLE_POS_LEVEL    ACCELCAL_VEHICLE_POS = 1
	ACCELCAL_VEHICLE_POS_LEFT     ACCELCAL_VEHICLE_POS = 2
	ACCELCAL_VEHICLE_POS_RIGHT    ACCELCAL_VEHICLE_POS = 3
	ACCELCAL_VEHICLE_POS_NOSEDOWN ACCELCAL_VEHICLE_POS = 4
	ACCELCAL_VEHICLE_POS_NOSEUP   ACCELCAL_VEHICLE_POS = 5
	ACCELCAL_VEHICLE_POS_BACK     ACCELCAL_VEHICLE_POS = 6
	ACCELCAL_VEHICLE_POS_SUCCESS  ACCELCAL_VEHICLE_POS = 16777215
	ACCELCAL_VEHICLE_POS_FAILED   ACCELCAL_VEHICLE_POS = 16777216
)

var labels_ACCELCAL_VEHICLE_POS = map[ACCELCAL_VEHICLE_POS]string{
	ACCELCAL_VEHICLE_POS_LEVEL:    "ACCELCAL_VEHICLE_POS_LEVEL",
	ACCELCAL_VEHICLE_POS_LEFT:     "ACCELCAL_VEHICLE_POS_LEFT",
	ACCELCAL_VEHICLE_POS_RIGHT:    "ACCELCAL_VEHICLE_POS_RIGHT",
	ACCELCAL_VEHICLE_POS_NOSEDOWN: "ACCELCAL_VEHICLE_POS_NOSEDOWN",
	ACCELCAL_VEHICLE_POS_NOSEUP:   "ACCELCAL_VEHICLE_POS_NOSEUP",
	ACCELCAL_VEHICLE_POS_BACK:     "ACCELCAL_VEHICLE_POS_BACK",
	ACCELCAL_VEHICLE_POS_SUCCESS:  "ACCELCAL_VEHICLE_POS_SUCCESS",
	ACCELCAL_VEHICLE_POS_FAILED:   "ACCELCAL_VEHICLE_POS_FAILED",
}

var values_ACCELCAL_VEHICLE_POS = map[string]ACCELCAL_VEHICLE_POS{
	"ACCELCAL_VEHICLE_POS_LEVEL":    ACCELCAL_VEHICLE_POS_LEVEL,
	"ACCELCAL_VEHICLE_POS_LEFT":     ACCELCAL_VEHICLE_POS_LEFT,
	"ACCELCAL_VEHICLE_POS_RIGHT":    ACCELCAL_VEHICLE_POS_RIGHT,
	"ACCELCAL_VEHICLE_POS_NOSEDOWN": ACCELCAL_VEHICLE_POS_NOSEDOWN,
	"ACCELCAL_VEHICLE_POS_NOSEUP":   ACCELCAL_VEHICLE_POS_NOSEUP,
	"ACCELCAL_VEHICLE_POS_BACK":     ACCELCAL_VEHICLE_POS_BACK,
	"ACCELCAL_VEHICLE_POS_SUCCESS":  ACCELCAL_VEHICLE_POS_SUCCESS,
	"ACCELCAL_VEHICLE_POS_FAILED":   ACCELCAL_VEHICLE_POS_FAILED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ACCELCAL_VEHICLE_POS) MarshalText() ([]byte, error) {
	if name, ok := labels_ACCELCAL_VEHICLE_POS[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ACCELCAL_VEHICLE_POS) UnmarshalText(text []byte) error {
	if value, ok := values_ACCELCAL_VEHICLE_POS[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = ACCELCAL_VEHICLE_POS(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e ACCELCAL_VEHICLE_POS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
