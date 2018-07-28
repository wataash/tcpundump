// Copyright (c) 2018, tcpundump authors
// All rights reserved.
// Licensed under BSD 2-Clause License.

package tcpundump

import (
	"os"
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

	scenario = "string \"foo\" doesn't contain hex dump, " +
		"consequently creates zero byte file"
	args = Args{FileWrite: "tmp.pcapng", Command: []string{"echo", "foo"}}
	if Tcpundump(args) != nil {
		t.Errorf("unexpected error in scenario: %q", scenario)
	}
	fi, err := os.Stat("tmp.pcapng")
	if err != nil {
		t.Errorf("unexpected error while os.Stat in scenario: %q", scenario)
	}
	if fi.Size() != 0 {
		t.Errorf(scenario)
	}

	scenario = "command's stderr goes tcpundump's one"
	args = Args{FileWrite: "tmp.pcapng",
		Command: []string{"sh", "-c", "echo foo 1>&2"}}
	// TODO: capture stderr
	if Tcpundump(args) != nil {
		t.Errorf("unexpected error in scenario: %q", scenario)
	}
}
