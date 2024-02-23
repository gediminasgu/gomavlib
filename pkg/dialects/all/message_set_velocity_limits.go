//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/development"
)

// Set temporary maximum limits for horizontal speed, vertical speed and yaw rate.
// The consumer must stream the current limits in VELOCITY_LIMITS at 1 Hz or more (when limits are being set).
// The consumer should latch the limits until a new limit is received or the mode is changed.
type MessageSetVelocityLimits = development.MessageSetVelocityLimits
