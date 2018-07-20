package tcpundump

import (
	"regexp"
	"fmt"
)

func parseDumpType(dumpType string) (dumpType, error) {
	switch dumpType {
	case "":
		return dtUnknown, nil
	case "cisco":
		return dtCisco, nil
	case "juniper":
		return dtJuniper, nil
	default:
		return dtUnknown, fmt.Errorf("%q: unknown dump type", dumpType)
	}
}

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
