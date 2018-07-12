package tcpundump

import (
	"regexp"
)

func inferType(line string) (dumpType, error) {
	// TODO improve performance: use last matched first

	for dumpType, regex := range regexInfer {
		matched, err := regexp.MatchString(regex, line)
		if err != nil {
			bugPanicUnknown(err, line)
		}
		if matched == true {
			return dumpType, nil
		}
	}

	// TODO: should be -1? what is the best practise?
	return -1, inferFailed
}
