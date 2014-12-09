// Package gore implement functions for detail errors
package gore

import (
	"fmt"
	"runtime"
)

// New returns an error with caller details
func New(text string) error {
	return newErr(text)
}

// Newf returns an error with formatting message and caller details
func Newf(text string, args ...interface{}) error {
	return newErr(fmt.Sprintf(text, args...))
}

func newErr(text string) error {
	pc, fn, line, _ := runtime.Caller(2)
	return &Err{
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fn,
		Line:     line,
		Msg:      text,
	}
}

// Err struct implement of error
type Err struct {
	FuncName string
	FileName string
	Line     int
	Msg      string
}

func (err *Err) Error() string {
	return err.Msg
}
