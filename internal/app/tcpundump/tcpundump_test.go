package tcpundump

import (
	"testing"
)

func TestTcpundump(t *testing.T) {
	var args Args
	var scenario string

	// test invalid arguments

	scenario = "can't specify both `-r` and `command` ([\"baz\"])"
	args = Args{FileRead: "foo", Command: []string{"baz"}}
	if Tcpundump(args).Error() != scenario {
		t.Errorf(scenario)
	}

	scenario = "open non-exist-fie: no such file or directory"
	args = Args{FileRead: "non-exist-fie"}
	if Tcpundump(args).Error() != scenario {
		t.Errorf(scenario)
	}

	// test cmd_reader.go

	args = Args{FileWrite: "tmp.pcapng", Command: []string{"echo", "foo"}}
	Tcpundump(args)
}
