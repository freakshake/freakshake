package xerror

import (
	"fmt"
)

// PanicIf will panic if err is not nil. It does nothing otherwise.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// Add adds context to the error.
// The result cannot be unwrapped to recover the original error.
// It does nothing when *errp == nil.
//
// Example:
//
//	defer xerror.Add(&err, "copy(%s, %s)", src, dst)
//
// See Wrap for an equivalent function that allows
// the result to be unwrapped.
func Add(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %v", fmt.Sprintf(format, args...), *errp)
	}
}

// Wrap adds context to the error and allows
// unwrapping the result to recover the original error.
//
// Example:
//
//	defer xerror.Wrap(&err, "copy(%s, %s)", src, dst)
//
// See Add for an equivalent function that does not allow
// the result to be unwrapped.
func Wrap(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}
