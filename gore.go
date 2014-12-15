// Package gore implement functions for detail errors
package gore

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
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
	return &Err{
		Msg:    text,
		Caller: newCallerInfo(2),
	}
}

// CallerInfo struct store info about caller
type CallerInfo struct {
	FuncName string
	FileName string
	Line     int
}

// String representation for CallerInfo object
func (ci *CallerInfo) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	buf.WriteString(filepath.Base(ci.FileName))
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(ci.Line))
	buf.WriteString(" ")
	buf.WriteString(filepath.Base(ci.FuncName))
	buf.WriteString("]")

	return buf.String()
}

func newCallerInfo(lvl int) *CallerInfo {
	pc, fn, line, _ := runtime.Caller(lvl + 1)
	return &CallerInfo{
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fn,
		Line:     line,
	}
}

// ContextElements slice store context value
type ContextElements []interface{}

// String representation for ContextElements object
func (ce ContextElements) String() string {
	var buf bytes.Buffer
	for _, v := range ce {
		buf.WriteString(fmt.Sprintf("%+v ", v))
	}

	return buf.String()
}

// Context struct store info about call context
type Context struct {
	Caller   *CallerInfo
	Elements ContextElements
}

// String representation of Context object
func (c *Context) String() string {
	return c.Caller.String() + " " + c.Elements.String()
}

// Err struct implement of error
type Err struct {
	Msg     string
	Caller  *CallerInfo
	Context []*Context
}

// Error return string for Err object
func (err *Err) Error() string {
	return err.Msg
}

func (err *Err) appendContext(args ...interface{}) {
	err.Context = append(err.Context, &Context{
		Caller:   newCallerInfo(2),
		Elements: args,
	})
}

// Append method append Context to given Err object
func Append(e error, args ...interface{}) {
	err, ok := e.(*Err)
	if !ok {
		return
	}
	err.appendContext(args...)
}
