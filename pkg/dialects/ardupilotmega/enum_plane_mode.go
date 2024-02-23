//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

// A mapping of plane flight modes for custom_mode field of heartbeat.
type PLANE_MODE uint32

const (
	PLANE_MODE_MANUAL        PLANE_MODE = 0
	PLANE_MODE_CIRCLE        PLANE_MODE = 1
	PLANE_MODE_STABILIZE     PLANE_MODE = 2
	PLANE_MODE_TRAINING      PLANE_MODE = 3
	PLANE_MODE_ACRO          PLANE_MODE = 4
	PLANE_MODE_FLY_BY_WIRE_A PLANE_MODE = 5
	PLANE_MODE_FLY_BY_WIRE_B PLANE_MODE = 6
	PLANE_MODE_CRUISE        PLANE_MODE = 7
	PLANE_MODE_AUTOTUNE      PLANE_MODE = 8
	PLANE_MODE_AUTO          PLANE_MODE = 10
	PLANE_MODE_RTL           PLANE_MODE = 11
	PLANE_MODE_LOITER        PLANE_MODE = 12
	PLANE_MODE_TAKEOFF       PLANE_MODE = 13
	PLANE_MODE_AVOID_ADSB    PLANE_MODE = 14
	PLANE_MODE_GUIDED        PLANE_MODE = 15
	PLANE_MODE_INITIALIZING  PLANE_MODE = 16
	PLANE_MODE_QSTABILIZE    PLANE_MODE = 17
	PLANE_MODE_QHOVER        PLANE_MODE = 18
	PLANE_MODE_QLOITER       PLANE_MODE = 19
	PLANE_MODE_QLAND         PLANE_MODE = 20
	PLANE_MODE_QRTL          PLANE_MODE = 21
	PLANE_MODE_QAUTOTUNE     PLANE_MODE = 22
	PLANE_MODE_QACRO         PLANE_MODE = 23
	PLANE_MODE_THERMAL       PLANE_MODE = 24
)

var labels_PLANE_MODE = map[PLANE_MODE]string{
	PLANE_MODE_MANUAL:        "PLANE_MODE_MANUAL",
	PLANE_MODE_CIRCLE:        "PLANE_MODE_CIRCLE",
	PLANE_MODE_STABILIZE:     "PLANE_MODE_STABILIZE",
	PLANE_MODE_TRAINING:      "PLANE_MODE_TRAINING",
	PLANE_MODE_ACRO:          "PLANE_MODE_ACRO",
	PLANE_MODE_FLY_BY_WIRE_A: "PLANE_MODE_FLY_BY_WIRE_A",
	PLANE_MODE_FLY_BY_WIRE_B: "PLANE_MODE_FLY_BY_WIRE_B",
	PLANE_MODE_CRUISE:        "PLANE_MODE_CRUISE",
	PLANE_MODE_AUTOTUNE:      "PLANE_MODE_AUTOTUNE",
	PLANE_MODE_AUTO:          "PLANE_MODE_AUTO",
	PLANE_MODE_RTL:           "PLANE_MODE_RTL",
	PLANE_MODE_LOITER:        "PLANE_MODE_LOITER",
	PLANE_MODE_TAKEOFF:       "PLANE_MODE_TAKEOFF",
	PLANE_MODE_AVOID_ADSB:    "PLANE_MODE_AVOID_ADSB",
	PLANE_MODE_GUIDED:        "PLANE_MODE_GUIDED",
	PLANE_MODE_INITIALIZING:  "PLANE_MODE_INITIALIZING",
	PLANE_MODE_QSTABILIZE:    "PLANE_MODE_QSTABILIZE",
	PLANE_MODE_QHOVER:        "PLANE_MODE_QHOVER",
	PLANE_MODE_QLOITER:       "PLANE_MODE_QLOITER",
	PLANE_MODE_QLAND:         "PLANE_MODE_QLAND",
	PLANE_MODE_QRTL:          "PLANE_MODE_QRTL",
	PLANE_MODE_QAUTOTUNE:     "PLANE_MODE_QAUTOTUNE",
	PLANE_MODE_QACRO:         "PLANE_MODE_QACRO",
	PLANE_MODE_THERMAL:       "PLANE_MODE_THERMAL",
}

var values_PLANE_MODE = map[string]PLANE_MODE{
	"PLANE_MODE_MANUAL":        PLANE_MODE_MANUAL,
	"PLANE_MODE_CIRCLE":        PLANE_MODE_CIRCLE,
	"PLANE_MODE_STABILIZE":     PLANE_MODE_STABILIZE,
	"PLANE_MODE_TRAINING":      PLANE_MODE_TRAINING,
	"PLANE_MODE_ACRO":          PLANE_MODE_ACRO,
	"PLANE_MODE_FLY_BY_WIRE_A": PLANE_MODE_FLY_BY_WIRE_A,
	"PLANE_MODE_FLY_BY_WIRE_B": PLANE_MODE_FLY_BY_WIRE_B,
	"PLANE_MODE_CRUISE":        PLANE_MODE_CRUISE,
	"PLANE_MODE_AUTOTUNE":      PLANE_MODE_AUTOTUNE,
	"PLANE_MODE_AUTO":          PLANE_MODE_AUTO,
	"PLANE_MODE_RTL":           PLANE_MODE_RTL,
	"PLANE_MODE_LOITER":        PLANE_MODE_LOITER,
	"PLANE_MODE_TAKEOFF":       PLANE_MODE_TAKEOFF,
	"PLANE_MODE_AVOID_ADSB":    PLANE_MODE_AVOID_ADSB,
	"PLANE_MODE_GUIDED":        PLANE_MODE_GUIDED,
	"PLANE_MODE_INITIALIZING":  PLANE_MODE_INITIALIZING,
	"PLANE_MODE_QSTABILIZE":    PLANE_MODE_QSTABILIZE,
	"PLANE_MODE_QHOVER":        PLANE_MODE_QHOVER,
	"PLANE_MODE_QLOITER":       PLANE_MODE_QLOITER,
	"PLANE_MODE_QLAND":         PLANE_MODE_QLAND,
	"PLANE_MODE_QRTL":          PLANE_MODE_QRTL,
	"PLANE_MODE_QAUTOTUNE":     PLANE_MODE_QAUTOTUNE,
	"PLANE_MODE_QACRO":         PLANE_MODE_QACRO,
	"PLANE_MODE_THERMAL":       PLANE_MODE_THERMAL,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PLANE_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_PLANE_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PLANE_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_PLANE_MODE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = PLANE_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e PLANE_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
