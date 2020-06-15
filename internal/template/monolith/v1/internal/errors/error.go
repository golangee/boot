package errors

// Try executes the given func and updates the error,
// but only if it has not been set yet.
func Try(f func() error, err *error) {
	newErr := f()

	if *err == nil {
		*err = newErr
	}
}
