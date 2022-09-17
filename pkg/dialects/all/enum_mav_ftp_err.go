//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"errors"
)

// MAV FTP error codes (https://mavlink.io/en/services/ftp.html)
type MAV_FTP_ERR uint32

const (
	// None: No error
	MAV_FTP_ERR_NONE MAV_FTP_ERR = 0
	// Fail: Unknown failure
	MAV_FTP_ERR_FAIL MAV_FTP_ERR = 1
	// FailErrno: Command failed, Err number sent back in PayloadHeader.data[1].
	// This is a file-system error number understood by the server operating system.
	MAV_FTP_ERR_FAILERRNO MAV_FTP_ERR = 2
	// InvalidDataSize: Payload size is invalid
	MAV_FTP_ERR_INVALIDDATASIZE MAV_FTP_ERR = 3
	// InvalidSession: Session is not currently open
	MAV_FTP_ERR_INVALIDSESSION MAV_FTP_ERR = 4
	// NoSessionsAvailable: All available sessions are already in use
	MAV_FTP_ERR_NOSESSIONSAVAILABLE MAV_FTP_ERR = 5
	// EOF: Offset past end of file for ListDirectory and ReadFile commands
	MAV_FTP_ERR_EOF MAV_FTP_ERR = 6
	// UnknownCommand: Unknown command / opcode
	MAV_FTP_ERR_UNKNOWNCOMMAND MAV_FTP_ERR = 7
	// FileExists: File/directory already exists
	MAV_FTP_ERR_FILEEXISTS MAV_FTP_ERR = 8
	// FileProtected: File/directory is write protected
	MAV_FTP_ERR_FILEPROTECTED MAV_FTP_ERR = 9
	// FileNotFound: File/directory not found
	MAV_FTP_ERR_FILENOTFOUND MAV_FTP_ERR = 10
)

var labels_MAV_FTP_ERR = map[MAV_FTP_ERR]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_FTP_ERR) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_FTP_ERR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_FTP_ERR = map[string]MAV_FTP_ERR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_FTP_ERR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_FTP_ERR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_FTP_ERR) String() string {
	if l, ok := labels_MAV_FTP_ERR[e]; ok {
		return l
	}
	return "invalid value"
}
