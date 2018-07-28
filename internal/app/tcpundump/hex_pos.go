// Copyright (c) 2018, tcpundump authors
// All rights reserved.
// Licensed under BSD 2-Clause License.

// TODO remane to col_byte.go

package tcpundump

import (
	"strings"
)

func isHex(char byte) bool {
	if ('0' <= char && char <= '9') || ('a' <= char && char <= 'f') ||
		('A' <= char && char <= 'F') {
		return true
	}
	return false
}

// Example:
//    012345678901234567890
//   " 0153 f2c1-9987  0 cc"
//     ^ ^  ^ ^  ^ ^     ^
//     1 3  6 8  1113    19
//   => [1 3 6 8 11 13 19]
// @param str len(str) must be <= colByteMax
func getColsByte(str string) []colByte {
	var cols []colByte

	if len(str) > colByteMax {
		bugPanicUnreachable(errUnreachable, str)
	}
	lenstr := colByte(len(str))

	//   " 0153 f2c1-9987  0 cc"
	//    FTFTFFTFTFFTFTFFFTFTF
	waitHex := false

	for i := colByte(0); i < lenstr; i++ {
		if isHex(str[i]) {
			if waitHex {
				cols = append(cols, i-1)
				waitHex = false
			} else {
				waitHex = true
			}
		} else {
			waitHex = false
		}
	}

	return cols
}

// TODO: audit
//   01234567890123456789012345678901234567890123456789012345678901234567890123
//   00000000001111111111222222222233333333334444444444555555555566666666667777
//
// Has full (16) bytes:
//                      _19  _24  _29  _34  _39  _44  _49  _54
//           0x0000:  1100 50b5 00aa 0021 4510 00a8 ce7d 4000  ..P....!E....}@.
//   ^0               ^17  ^22  ^27  ^32  ^37  ^42  ^47  ^52
//   => [17 19 22 24 27 29 32 34 37 39 42 44 47 49 52 54]
//
// Has the end of the packet which has a length less than 16 bytes:
//           0x0060:  050b 0d                                  foobar
//   ^0               ^ ^  ^ 17, 19, 22
//   => [17 19 22]
//
// @error nil, errUnreachable
func lineToColsByteSeil(line string) ([]colByte, error) {
	if len(line) >= 56 {
		// has ascii dump, remove it
		line = line[:56]
		line = strings.TrimRight(line, " ")
	}

	cols := []colByte{
		17, 19, 22, 24, 27, 29, 32, 34, 37, 39, 42, 44, 47, 49, 52, 54}

	if len(line) > colByteMax {
		return []colByte{}, errUnreachable
	}
	length := colByte(len(line))

	for i, col := range cols {
		if length == col+2 {
			return cols[:i+1], nil
		}
	}

	return []colByte{}, errUnreachable
}

func lineToHexPositionsTcpdumpX(line string) ([]colByte, error) {
	bugPanic(nil, "TODO")
	return []colByte{}, nil
}

func lineToColsByte(line string) ([]colByte, error) {
	hexType, err := inferType(line)

	switch err {
	case nil:
		break
	case errInferFailed:
		return []colByte{}, errInferFailed
	default:
		bugPanicUnknown(err, line)
	}

	f, ok := funcHexPos[hexType]
	if ok != true {
		bugPanicUnreachable(err, line)
	}

	return f(line)
}
