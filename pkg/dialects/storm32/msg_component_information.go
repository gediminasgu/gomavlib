//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Component information message, which may be requested using MAV_CMD_REQUEST_MESSAGE.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// CRC32 of the general metadata file (general_metadata_uri).
	GeneralMetadataFileCrc uint32
	// MAVLink FTP URI for the general metadata file (COMP_METADATA_TYPE_GENERAL), which may be compressed with xz. The file contains general component metadata, and may contain URI links for additional metadata (see COMP_METADATA_TYPE). The information is static from boot, and may be generated at compile time. The string needs to be zero terminated.
	GeneralMetadataUri string `mavlen:"100"`
	// CRC32 of peripherals metadata file (peripherals_metadata_uri).
	PeripheralsMetadataFileCrc uint32
	// (Optional) MAVLink FTP URI for the peripherals metadata file (COMP_METADATA_TYPE_PERIPHERALS), which may be compressed with xz. This contains data about "attached components" such as UAVCAN nodes. The peripherals are in a separate file because the information must be generated dynamically at runtime. The string needs to be zero terminated.
	PeripheralsMetadataUri string `mavlen:"100"`
}

// GetID implements the message.Message interface.
func (*MessageComponentInformation) GetID() uint32 {
	return 395
}
