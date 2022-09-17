//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Set the mission item with sequence number seq as the current item and emit MISSION_CURRENT (whether or not the mission number changed).
// If a mission is currently being executed, the system will continue to this new mission item on the shortest path, skipping any intermediate mission items.
// Note that mission jump repeat counters are not reset (see MAV_CMD_DO_JUMP param2).
// This message may trigger a mission state-machine change on some systems: for example from MISSION_STATE_NOT_STARTED or MISSION_STATE_PAUSED to MISSION_STATE_ACTIVE.
// If the system is in mission mode, on those systems this command might therefore start, restart or resume the mission.
// If the system is not in mission mode this message must not trigger a switch to mission mode.
type MessageMissionSetCurrent struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
}

// GetID implements the message.Message interface.
func (*MessageMissionSetCurrent) GetID() uint32 {
	return 41
}
