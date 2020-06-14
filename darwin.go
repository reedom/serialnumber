// +build darwin

package serialnumber

import (
	"os/exec"
)

func Read() (string, error) {
	out, err := exec.Command("/usr/sbin/ioreg", "-l").Output()
	if err != nil {
		return "", newError(err)
	}
	return matchDarwin(out)
}
