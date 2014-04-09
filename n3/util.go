package n3

// This is a place for random utitlites to go

// this returns the first error, or no errors
func checkAnyError(errs ...error) (err error) {
	for _, err = range errs {
		if err != nil {
			return
		}
	}

	return
}
