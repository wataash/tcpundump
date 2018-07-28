// Copyright (c) 2018, tcpundump authors
// All rights reserved.
// Licensed under BSD 2-Clause License.

package tcpundump

import "fmt"

func ExampleUndumpByte() {
	fmt.Println(undumpByte("00"))
	fmt.Println(undumpByte("0f"))
	fmt.Println(undumpByte("ff"))
	fmt.Println(undumpByte("foo"))
	// Output:
	// 0 <nil>
	// 15 <nil>
	// 255 <nil>
	// 0 unreachable
}

func ExampleUndump() {
	fmt.Println(undump("  01 23 ab f ", []colByte{2, 5, 8}))
	fmt.Println(undump("", []colByte{}))
	fmt.Println(undump("foo", []colByte{1, 2, 3, 4, 5}))
	// Output:
	// [1 35 171] <nil>
	// [] <nil>
	// [] unreachable
}
