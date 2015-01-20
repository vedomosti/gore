package gore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	assert := assert.New(t)
	err := New("hello")
	assert.Equal(err.Error(), "hello")

	err = Newf("hello: %s", "world")
	assert.Equal(err.Error(), "hello: world")

	caller := err.(*Err).Caller
	assert.Equal(caller.ShortFileName(), "gore_test.go:14")
	assert.Equal(caller.ShortFuncName(), "gore.TestErr")

	Append(err, "foo")
	assert.Equal(err.(*Err).Context[0].String(), "foo")

	Appendf(err, "foo: %s", "bar")
	assert.Equal(err.(*Err).Context[1].String(), "foo: bar")
	assert.NotNil(err.(*Err).Context[0].Caller)
}
