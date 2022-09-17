// Package minimal contains the minimal dialect.
//
//autogenerated:yes
package minimal

import (
	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/message"
)

// Dialect contains the dialect definition.
var Dialect = dial

// dial is not exposed directly in order not to display it in godoc.
var dial = &dialect.Dialect{
	Version: 3,
	Messages: []message.Message{
		// minimal.xml
		&MessageHeartbeat{},
		&MessageProtocolVersion{},
	},
}
