package xerror

import (
	"errors"
	"runtime"
)

// WrapStack is like Wrap, but adds a stack trace if there isn't one already.
func WrapStack(errp *error, format string, args ...any) {
	if *errp != nil {
		if se := (*StackError)(nil); !errors.As(*errp, &se) {
			*errp = NewStackError(*errp)
		}
		Wrap(errp, format, args...)
	}
}

// StackError wraps an error and adds a stack trace.
type StackError struct {
	Stack []byte
	err   error
}

// NewStackError returns a StackError, capturing a stack trace.
func NewStackError(err error) *StackError {
	var buf [16 * 1024]byte
	n := runtime.Stack(buf[:], false)
	return &StackError{
		err:   err,
		Stack: buf[:n],
	}
}

func (e *StackError) Error() string {
	return e.err.Error() // ignore the stack
}

func (e *StackError) Unwrap() error {
	return e.err
}
