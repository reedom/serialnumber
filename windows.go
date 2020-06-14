// +build windows

package serialnumber

func Read() (string, error) {
	out, err := exec.Command("wmic", "bios get serialnumber").Output()
	if err != nil {
		return "", newError(err)
	}
	return matchWindows(out)
}
