package serialnumber

import (
	"fmt"
	"regexp"
)

func matchDarwin(out []byte) (string, error) {
	re := regexp.MustCompile(`"IOPlatformSerialNumber" = "(.*?)"`)
	if m := re.FindSubmatch(out); m != nil {
		return string(m[1]), nil
	}

	return "", newError(fmt.Errorf("ioreg output format has been changed"))
}

func matchWindows(out []byte) (string, error) {
	re := regexp.MustCompile(`(?i)SerialNumber[\s]+(\S+)`)
	if m := re.FindSubmatch(out); m != nil {
		return string(m[1]), nil
	}

	return "", newError(fmt.Errorf("wmic output format has been changed"))
}

