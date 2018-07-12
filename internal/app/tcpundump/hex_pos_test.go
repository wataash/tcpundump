package tcpundump

import (
	"fmt"
)

func ExampleIsHex_digit() {
	fmt.Println(isHex('0'))
	fmt.Println(isHex('9'))
	// Output:
	// true
	// true
}

func ExampleIsHex_lower() {
	fmt.Println(isHex('a'))
	fmt.Println(isHex('f'))
	fmt.Println(isHex('g'))
	fmt.Println(isHex('x'))
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleIsHex_upper() {
	fmt.Println(isHex('A'))
	fmt.Println(isHex('F'))
	fmt.Println(isHex('G'))
	// Output:
	// true
	// true
	// false
}

func ExampleIsHex_etc() {
	fmt.Println(isHex('-'))
	fmt.Println(isHex('$'))
	// Output:
	// false
	// false
}

func ExampleGetColsByte() {
	fmt.Println(getColsByte("f"))
	fmt.Println(getColsByte("ff"))
	fmt.Println(getColsByte(" 0153 f2c1-9987  0 cc"))
	// Output:
	// []
	// [0]
	// [1 3 6 8 11 13 19]
}

func ExampleHexPositions_no_hex() {
	fmt.Println(getColsByte(""))
	fmt.Println(getColsByte("X"))
	fmt.Println(getColsByte("XXX"))
	// Output:
	// []
	// []
	// []
}

func ExampleColsByteSeil_full() {
	line := "        0x0000:  1100 50b5 00aa 0021 4510 00a8 ce7d 4000" +
		"  ..P....!E....}@."
	positions, err := lineToColsByteSeil(line)
	fmt.Println(positions)
	fmt.Println(err)
	// Output:
	// [17 19 22 24 27 29 32 34 37 39 42 44 47 49 52 54]
	// <nil>
}

func ExampleColsByteSeil_end() {
	line := "        0x0060:  050b 0d                                " +
		"  foobar"
	positions, err := lineToColsByteSeil(line)
	fmt.Println(positions)
	fmt.Println(err)
	// Output:
	// [17 19 22]
	// <nil>
}

func ExampleLineToColsByte() {
	// SEIL
	const line1 = "        0x0000:  1100 50b5 00aa 0021 4510 00a8 ce7d 4000" +
		"  ..P....!E....}@."
	positions, err := lineToColsByte(line1)
	fmt.Println(positions)
	fmt.Println(err)
	// Output:
	// [17 19 22 24 27 29 32 34 37 39 42 44 47 49 52 54]
	// <nil>
}
