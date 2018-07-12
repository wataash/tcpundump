package tcpundump

import "strconv"

// "00" -> 0
// "0f" -> 15
// "ff" -> 255
// @error unreachable, strconv.ParseInt()'s error
func undumpByte(dump string) (byte, error) {
	if len(dump) != 2 || !isHex(dump[0]) || !isHex(dump[1]) {
		return 0, unreachable
	}

	b, err := strconv.ParseUint(dump, 16, 8)
	return byte(b), err
}

// @example
//   "  01 23 ab f "
//      ^2 ^5 ^8
//   => [0x01, 0x23, 0xab]
// @error unreachable, strconv.ParseInt()'s error
func undump(line string, colsByte []colByte) ([]byte, error) {
	undumped := make([]byte, 0, len(line)/2)

	for i, col := range colsByte {
		_, _ = i, col
		b, err := undumpByte(line[col : col+2])
		if err != nil {
			return []byte{}, err
		}
		undumped = append(undumped, b)
	}
	return undumped, nil
}
