package serialnumber

import (
	"testing"
)

func TestDarwin(t *testing.T) {
	out := []byte(`
    |   "manufacturer" = <"Apple Inc.">
    |   "IOPlatformSerialNumber" = "A01B02C3DEFG"
    |   "system-type" = <02>
`)
	expected := "A01B02C3DEFG"

	actual, err := matchDarwin(out)
	if err != nil {
		t.Errorf("err != nil")
		return
	}
	if actual != expected {
		t.Errorf("got %#v, want %#v", actual, expected)
	}
}

func TestDarwinUnmatch(t *testing.T) {
	_, err := matchDarwin([]byte(""))
	if err == nil {
		t.Errorf("err == nil")
		return
	}
}

func TestWindows(t *testing.T) {
	out := []byte(`
SerialNumber
1234567890
`)
	expected := "1234567890"

	actual, err := matchWindows(out)
	if err != nil {
		t.Errorf("err != nil")
		return
	}
	if actual != expected {
		t.Errorf("got %#v, want %#v", actual, expected)
	}
}

func TestWindowsUnmatch(t *testing.T) {
	_, err := matchWindows([]byte(""))
	if err == nil {
		t.Errorf("err == nil")
		return
	}
}

