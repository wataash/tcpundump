package tcpundump

import (
	"fmt"
	"testing"
)

func TestTcpundump(t *testing.T) {
	var args Args
	var scenario string

	// test invalid arguments

	scenario = "can't specify both `-r` and `command` ([\"baz\"])"
	args = Args{FileRead: "foo", Command: []string{"baz"}}
	if Tcpundump(args).Error() != scenario {
		fmt.Errorf(scenario)
	}

	scenario = "-w must be specified if `command` ([\"baz\"]) is given"
	args = Args{Command: []string{"baz"}}
	if Tcpundump(args).Error() != scenario {
		fmt.Errorf(scenario)
	}

	scenario = "open non-exist-fie: no such file or directory"
	args = Args{FileRead: "non-exist-fie"}
	if Tcpundump(args).Error() != scenario {
		fmt.Errorf(scenario)
	}

	// test read.go

	args = Args{FileWrite: "tmp.pcapng", Command: []string{"echo", "foo"}}
	Tcpundump(args)
	fmt.Println("foo")
}
