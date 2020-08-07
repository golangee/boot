package errors

// Check executes the given func and updates the error,
// but only if it has not been set yet. This should be
// used to handle defer and io.Closer correctly.
//
// Example
//  func do(fname string) (err error) {
//    file, err := os.Open(fname)
//    if err != nil{
//       return err
//    }
//
//    defer errors.Check(r.Close, &err)
//
//    // do stuff with file
//  }
func Check(f func() error, err *error) {
	newErr := f()

	if *err == nil {
		*err = newErr
	}
}
