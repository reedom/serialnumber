package serialnumber

func newError(err error) ReadError {
	return ReadError{err}
}

type ReadError struct {
	err error
}

func (w ReadError) Error() string {
	return "failed to read a serial number"
}

func (w ReadError) Wrap() error {
	return w.err
}
