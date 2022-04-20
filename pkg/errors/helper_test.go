package errors

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestIsOneOf(t *testing.T) {
	assert := require.New(t)
	err1 := New("err-1")
	err2 := New("err-2")
	err3 := Wrap(err2, "err-3")

	// pure error
	assert.True(IsOneOf(err1, err1))
	assert.False(IsOneOf(err1, err2))

	// error with wrap
	assert.True(IsOneOf(err3, err2))
	assert.False(IsOneOf(err3, err1))

	// any targets
	assert.True(IsOneOf(err1, err2, err1))
	assert.True(IsOneOf(err3, err2, err1))
	assert.False(IsOneOf(err1, err2, err3))
}

func TestCause(t *testing.T) {
	err1 := errors.New("err-1")
	err2 := errors.Wrap(err1, "err-2")
	err3 := errors.Wrap(err2, "err-3")
	err4 := errors.New("err-4")

	require.Equal(t, err1, errors.Cause(err3))
	require.Equal(t, err1, errors.Cause(err2))
	require.NotEqual(t, err1, errors.Cause(err4))
}
