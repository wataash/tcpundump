package tcpundump

import "fmt"

func ExampleTcpundump() {
	var args Args

	args = Args{FileRead: "foo", Command: []string{"baz"}}
	fmt.Println(Tcpundump(args))

	args = Args{Command: []string{"baz"}}
	fmt.Println(Tcpundump(args))

	args = Args{FileRead: "non-exist-fie"}
	fmt.Println(Tcpundump(args))

	// Output:
	// can't specify both `-r` and `command` (["baz"])
	// -w must be specified if `command` (["baz"]) is given
	// open non-exist-fie: no such file or directory
}
