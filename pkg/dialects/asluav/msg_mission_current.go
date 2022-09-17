//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Message that announces the sequence number of the current target mission item (that the system will fly towards/execute when the mission is running).
// This message should be streamed all the time (nominally at 1Hz).
// This message should be emitted following a call to MAV_CMD_DO_SET_MISSION_CURRENT or SET_MISSION_CURRENT.
type MessageMissionCurrent struct {
	// Sequence
	Seq uint16
	// Total number of mission items. 0: Not supported, UINT16_MAX if no mission is present on the vehicle.
	Total uint16 `mavext:"true"`
	// Mission state machine state. MISSION_STATE_UNKNOWN if state reporting not supported.
	MissionState MISSION_STATE `mavenum:"uint8" mavext:"true"`
	// Vehicle is in a mode that can execute mission items or suspended. 0: Unknown, 1: In mission mode, 2: Suspended (not in mission mode).
	MissionMode uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageMissionCurrent) GetID() uint32 {
	return 42
}
