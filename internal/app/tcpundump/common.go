package tcpundump

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// ----------------------------------------------------------------------------
// type

type dumpType float64

// Assuming hex dump always appears as a two-digit bytes:
//
//   01 23 45 67 89 ab c
//   ^^ ^^ ^^ ^^ ^^ ^^ ^
//   0x01, 0x23, 0x45, 0x67, 0x89, 0xab (, "c" is not interpreted as a byte)
//
// The column of a byte is the column of the most significant bit (the second
// digit):
//
//   01 23 45 67 89 ab c
//   ^  ^  ^  ^  ^  ^
//   0  3  6  9  12 15
type colByte uint8

const colByteMax = math.MaxUint8

var inferFailed = errors.New("infer failed")
var invalidInput = errors.New("invalid input")
var unreachable = errors.New("unreachable")

// ----------------------------------------------------------------------------
// const

// TODO: move to dump_type.go?
const (
	_ dumpType = iota
	dtCisco
	dtJuniper
	dtJuniper2
	dtSeil
	dtTcpdump_old_TODO
	dtTcpdumpX  // -x
	dtTcpdumpXX // -X
	dtUnknown
)

// want to be a const
var regexInfer = map[dumpType]string{
	// TODO audit
	// (8 SP)0x0000:  1100 50b5 00aa 0021 4510 00a8 ce7d 4000  ..P....!E....}@.
	dtSeil:     ` {8}0x\d{3}0:  [[:xdigit:]]`,
	dtTcpdumpX: `\t0x\d{3}0:  [[:xdigit:]]`,
}

var funcHexPos = map[dumpType]func(line string) ([]colByte, error){
	dtSeil:     lineToColsByteSeil,
	dtTcpdumpX: lineToHexPositionsTcpdumpX,
}

// ----------------------------------------------------------------------------
// func

// TODO doc: Must be called in functions which have information like line
// Don't execute me in like undumpByte().
func bugPanic(err error, format string, a ...interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintf(os.Stderr, "tcpundump: BUG: "+format+"\n", a...)
	s := "Please report issue to https://github.com/wataash/tcpundump/issues "
	s += "if this error is reproducible."
	fmt.Fprintln(os.Stderr, s)
	panic(nil)
}

func bugPanicUnknown(err error, line string) {
	bugPanic(err, "Unknown error while line a line: %s", line)
}

func bugPanicUnreachable(err error, line string) {
	if err == unreachable {
		// Suppress a redundant error message
		err = nil
	}
	bugPanic(err, "Unreachable code executed while line a line: %s", line)
}
