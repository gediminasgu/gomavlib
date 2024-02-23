//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Set temporary maximum limits for horizontal speed, vertical speed and yaw rate.
// The consumer must stream the current limits in VELOCITY_LIMITS at 1 Hz or more (when limits are being set).
// The consumer should latch the limits until a new limit is received or the mode is changed.
type MessageSetVelocityLimits struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Limit for horizontal movement in MAV_FRAME_LOCAL_NED. NaN: Field not used (ignore)
	HorizontalSpeedLimit float32
	// Limit for vertical movement in MAV_FRAME_LOCAL_NED. NaN: Field not used (ignore)
	VerticalSpeedLimit float32
	// Limit for vehicle turn rate around its yaw axis. NaN: Field not used (ignore)
	YawRateLimit float32
}

// GetID implements the message.Message interface.
func (*MessageSetVelocityLimits) GetID() uint32 {
	return 354
}
