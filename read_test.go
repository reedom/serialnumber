package serialnumber

import (
	"testing"
)

func TestRead(t *testing.T) {
	sn, err := Read()
	if err != nil {
		t.Errorf("err != nil")
		return
	}
	if len(sn) < 10 {
		t.Errorf("obtained serial number seems wrong: %#v", sn)
		return
	}
}
