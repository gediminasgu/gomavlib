//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"fmt"
	"strconv"
	"strings"
)

// Battery status flags for fault, health and state indication.
type MAV_BATTERY_STATUS_FLAGS uint32

const (
	// The battery is not ready to use (fly).
	// Set if the battery has faults or other conditions that make it unsafe to fly with.
	// Note: It will be the logical OR of other status bits (chosen by the manufacturer/integrator).
	MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE MAV_BATTERY_STATUS_FLAGS = 1
	// Battery is charging.
	MAV_BATTERY_STATUS_FLAGS_CHARGING MAV_BATTERY_STATUS_FLAGS = 2
	// Battery is cell balancing (during charging).
	// Not ready to use (MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE may be set).
	MAV_BATTERY_STATUS_FLAGS_CELL_BALANCING MAV_BATTERY_STATUS_FLAGS = 4
	// Battery cells are not balanced.
	// Not ready to use.
	MAV_BATTERY_STATUS_FLAGS_FAULT_CELL_IMBALANCE MAV_BATTERY_STATUS_FLAGS = 8
	// Battery is auto discharging (towards storage level).
	// Not ready to use (MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE would be set).
	MAV_BATTERY_STATUS_FLAGS_AUTO_DISCHARGING MAV_BATTERY_STATUS_FLAGS = 16
	// Battery requires service (not safe to fly).
	// This is set at vendor discretion.
	// It is likely to be set for most faults, and may also be set according to a maintenance schedule (such as age, or number of recharge cycles, etc.).
	MAV_BATTERY_STATUS_FLAGS_REQUIRES_SERVICE MAV_BATTERY_STATUS_FLAGS = 32
	// Battery is faulty and cannot be repaired (not safe to fly).
	// This is set at vendor discretion.
	// The battery should be disposed of safely.
	MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY MAV_BATTERY_STATUS_FLAGS = 64
	// Automatic battery protection monitoring is enabled.
	// When enabled, the system will monitor for certain kinds of faults, such as cells being over-voltage.
	// If a fault is triggered then and protections are enabled then a safety fault (MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM) will be set and power from the battery will be stopped.
	// Note that battery protection monitoring should only be enabled when the vehicle is landed. Once the vehicle is armed, or starts moving, the protections should be disabled to prevent false positives from disabling the output.
	MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED MAV_BATTERY_STATUS_FLAGS = 128
	// The battery fault protection system had detected a fault and cut all power from the battery.
	// This will only trigger if MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED is set.
	// Other faults like MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT may also be set, indicating the cause of the protection fault.
	MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM MAV_BATTERY_STATUS_FLAGS = 256
	// One or more cells are above their maximum voltage rating.
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT MAV_BATTERY_STATUS_FLAGS = 512
	// One or more cells are below their minimum voltage rating.
	// A battery that had deep-discharged might be irrepairably damaged, and set both MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT and MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY.
	MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT MAV_BATTERY_STATUS_FLAGS = 1024
	// Over-temperature fault.
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_TEMPERATURE MAV_BATTERY_STATUS_FLAGS = 2048
	// Under-temperature fault.
	MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_TEMPERATURE MAV_BATTERY_STATUS_FLAGS = 4096
	// Over-current fault.
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_CURRENT MAV_BATTERY_STATUS_FLAGS = 8192
	// Short circuit event detected.
	// The battery may or may not be safe to use (check other flags).
	MAV_BATTERY_STATUS_FLAGS_FAULT_SHORT_CIRCUIT MAV_BATTERY_STATUS_FLAGS = 16384
	// Voltage not compatible with power rail voltage (batteries on same power rail should have similar voltage).
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_VOLTAGE MAV_BATTERY_STATUS_FLAGS = 32768
	// Battery firmware is not compatible with current autopilot firmware.
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_FIRMWARE MAV_BATTERY_STATUS_FLAGS = 65536
	// Battery is not compatible due to cell configuration (e.g. 5s1p when vehicle requires 6s).
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION MAV_BATTERY_STATUS_FLAGS = 131072
	// Battery capacity_consumed and capacity_remaining values are relative to a full battery (they sum to the total capacity of the battery).
	// This flag would be set for a smart battery that can accurately determine its remaining charge across vehicle reboots and discharge/recharge cycles.
	// If unset the capacity_consumed indicates the consumption since vehicle power-on, as measured using a power monitor. The capacity_remaining, if provided, indicates the estimated remaining capacity on the assumption that the battery was full on vehicle boot.
	// If unset a GCS is recommended to advise that users fully charge the battery on power on.
	MAV_BATTERY_STATUS_FLAGS_CAPACITY_RELATIVE_TO_FULL MAV_BATTERY_STATUS_FLAGS = 262144
	// Reserved (not used). If set, this will indicate that an additional status field exists for higher status values.
	MAV_BATTERY_STATUS_FLAGS_EXTENDED MAV_BATTERY_STATUS_FLAGS = 4294967295
)

var labels_MAV_BATTERY_STATUS_FLAGS = map[MAV_BATTERY_STATUS_FLAGS]string{
	MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE:                       "MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE",
	MAV_BATTERY_STATUS_FLAGS_CHARGING:                               "MAV_BATTERY_STATUS_FLAGS_CHARGING",
	MAV_BATTERY_STATUS_FLAGS_CELL_BALANCING:                         "MAV_BATTERY_STATUS_FLAGS_CELL_BALANCING",
	MAV_BATTERY_STATUS_FLAGS_FAULT_CELL_IMBALANCE:                   "MAV_BATTERY_STATUS_FLAGS_FAULT_CELL_IMBALANCE",
	MAV_BATTERY_STATUS_FLAGS_AUTO_DISCHARGING:                       "MAV_BATTERY_STATUS_FLAGS_AUTO_DISCHARGING",
	MAV_BATTERY_STATUS_FLAGS_REQUIRES_SERVICE:                       "MAV_BATTERY_STATUS_FLAGS_REQUIRES_SERVICE",
	MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY:                            "MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY",
	MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED:                    "MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED",
	MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM:                "MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM",
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT:                        "MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT",
	MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT:                       "MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT",
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_TEMPERATURE:                 "MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_TEMPERATURE",
	MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_TEMPERATURE:                "MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_TEMPERATURE",
	MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_CURRENT:                     "MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_CURRENT",
	MAV_BATTERY_STATUS_FLAGS_FAULT_SHORT_CIRCUIT:                    "MAV_BATTERY_STATUS_FLAGS_FAULT_SHORT_CIRCUIT",
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_VOLTAGE:             "MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_VOLTAGE",
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_FIRMWARE:            "MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_FIRMWARE",
	MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION: "MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION",
	MAV_BATTERY_STATUS_FLAGS_CAPACITY_RELATIVE_TO_FULL:              "MAV_BATTERY_STATUS_FLAGS_CAPACITY_RELATIVE_TO_FULL",
	MAV_BATTERY_STATUS_FLAGS_EXTENDED:                               "MAV_BATTERY_STATUS_FLAGS_EXTENDED",
}

var values_MAV_BATTERY_STATUS_FLAGS = map[string]MAV_BATTERY_STATUS_FLAGS{
	"MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE":                       MAV_BATTERY_STATUS_FLAGS_NOT_READY_TO_USE,
	"MAV_BATTERY_STATUS_FLAGS_CHARGING":                               MAV_BATTERY_STATUS_FLAGS_CHARGING,
	"MAV_BATTERY_STATUS_FLAGS_CELL_BALANCING":                         MAV_BATTERY_STATUS_FLAGS_CELL_BALANCING,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_CELL_IMBALANCE":                   MAV_BATTERY_STATUS_FLAGS_FAULT_CELL_IMBALANCE,
	"MAV_BATTERY_STATUS_FLAGS_AUTO_DISCHARGING":                       MAV_BATTERY_STATUS_FLAGS_AUTO_DISCHARGING,
	"MAV_BATTERY_STATUS_FLAGS_REQUIRES_SERVICE":                       MAV_BATTERY_STATUS_FLAGS_REQUIRES_SERVICE,
	"MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY":                            MAV_BATTERY_STATUS_FLAGS_BAD_BATTERY,
	"MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED":                    MAV_BATTERY_STATUS_FLAGS_PROTECTIONS_ENABLED,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM":                MAV_BATTERY_STATUS_FLAGS_FAULT_PROTECTION_SYSTEM,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT":                        MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_VOLT,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT":                       MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_VOLT,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_TEMPERATURE":                 MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_TEMPERATURE,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_TEMPERATURE":                MAV_BATTERY_STATUS_FLAGS_FAULT_UNDER_TEMPERATURE,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_CURRENT":                     MAV_BATTERY_STATUS_FLAGS_FAULT_OVER_CURRENT,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_SHORT_CIRCUIT":                    MAV_BATTERY_STATUS_FLAGS_FAULT_SHORT_CIRCUIT,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_VOLTAGE":             MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_VOLTAGE,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_FIRMWARE":            MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_FIRMWARE,
	"MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION": MAV_BATTERY_STATUS_FLAGS_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION,
	"MAV_BATTERY_STATUS_FLAGS_CAPACITY_RELATIVE_TO_FULL":              MAV_BATTERY_STATUS_FLAGS_CAPACITY_RELATIVE_TO_FULL,
	"MAV_BATTERY_STATUS_FLAGS_EXTENDED":                               MAV_BATTERY_STATUS_FLAGS_EXTENDED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_STATUS_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 20; i++ {
		mask := MAV_BATTERY_STATUS_FLAGS(1 << i)
		if e&mask == mask {
			names = append(names, labels_MAV_BATTERY_STATUS_FLAGS[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_STATUS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_BATTERY_STATUS_FLAGS
	for _, label := range labels {
		if value, ok := values_MAV_BATTERY_STATUS_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= MAV_BATTERY_STATUS_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_STATUS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
