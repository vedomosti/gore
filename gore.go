// Package gore implement functions for detail errors
package gore

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
)

// Err struct implement of error
type Err struct {
	Msg     string
	Caller  *Caller
	Context []*Context
}

// New returns an error with caller details
func New(text string) error {
	return newErr(text)
}

// Newf returns an error with formatting message and caller details
func Newf(text string, args ...interface{}) error {
	return newErr(fmt.Sprintf(text, args...))
}

func newErr(text string) error {
	return &Err{
		Msg:    text,
		Caller: NewCaller(2),
	}
}

// Error return string for Err object
func (err *Err) Error() string {
	return err.Msg
}

// Caller struct store info about caller
type Caller struct {
	FuncName string
	FileName string
	Line     int
}

func NewCaller(lvl int) *Caller {
	pc, fn, line, _ := runtime.Caller(lvl + 1)
	return &Caller{
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fn,
		Line:     line,
	}
}

func (caller *Caller) ShortFileName() string {
	return filepath.Base(caller.FileName) + ":" + strconv.Itoa(caller.Line)
}

func (caller *Caller) ShortFuncName() string {
	return filepath.Base(caller.FuncName)
}

// Context struct store info about call context
type Context struct {
	Caller *Caller
	Msg    string
}

// String representation of Context object
func (c *Context) String() string {
	return c.Msg
}

// Append method append Context to given Err object
func Append(e error, args ...interface{}) {
	err, ok := e.(*Err)
	if !ok {
		return
	}
	appendContext(err, fmt.Sprint(args...))
}

func Appendf(e error, format string, args ...interface{}) {
	err, ok := e.(*Err)
	if !ok {
		return
	}
	appendContext(err, fmt.Sprintf(format, args...))
}

func appendContext(err *Err, msg string) {
	err.Context = append(err.Context, &Context{
		Caller: NewCaller(2),
		Msg:    msg,
	})
}
