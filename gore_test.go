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

	caller := err.Caller
	assert.Equal(caller.ShortFileName(), "gore_test.go:14")
	assert.Equal(caller.ShortFuncName(), "gore.TestErr")

	Append(err, "foo")
	assert.Equal(err.Context[0].String(), "foo")

	gerr := Appendf(err, "foo: %s", "bar")
	assert.Equal(err.Context[1].String(), "foo: bar")
	assert.NotNil(err.Context[0].Caller)

	gerr.Append("append1").Appendf("append%d", 2)
	assert.Equal(gerr.Context[2].String(), "append1")
	assert.Equal(gerr.Context[3].String(), "append2")
}
