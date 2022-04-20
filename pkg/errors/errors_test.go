package errors

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoggerForCollectAndAppend(t *testing.T) {
	assert := require.New(t)

	read, cancel := newReadFunc(t)
	defer cancel()

	collector := NewErrorCollector("multi error")
	assert.EqualError(collector, "")

	// skiped
	collector.Append(nil)
	collector.Collect(nil, "error description")

	// one error
	collector.Collect(New("first fail"), "prefix #1")
	assert.EqualError(collector, "multi error: prefix #1: first fail")

	// two errors
	collector.Collect(New("second fail"), "prefix #2")
	assert.EqualError(collector, "multi error: prefix #1: first fail; prefix #2: second fail")

	// three errors
	collector.Append(New("third fail"))
	assert.EqualError(collector, "multi error: prefix #1: first fail; prefix #2: second fail; third fail")

	checkLogMessages(
		t,
		read(),
		"errors_test.go",
		"prefix #1: first fail",
		"prefix #2: second fail",
		"third fail",
	)
}

func TestLoggerForCollectf(t *testing.T) {
	assert := require.New(t)

	read, cancel := newReadFunc(t)
	defer cancel()

	collector := NewErrorCollector("multi error")
	assert.EqualError(collector, "")

	// skiped
	collector.Collectf(nil, "error description")

	// one error
	collector.Collectf(New("first fail"), "prefix #1 '%s,%s'", "arg1", "arg2")
	assert.EqualError(collector, "multi error: prefix #1 'arg1,arg2': first fail")

	// two errors
	collector.Collectf(New("second fail"), "prefix #2")
	assert.EqualError(collector, "multi error: prefix #1 'arg1,arg2': first fail; prefix #2: second fail")

	checkLogMessages(
		t,
		read(),
		"errors_test.go",
		"prefix #1 'arg1,arg2': first fail",
		"prefix #2: second fail",
	)
}

func TestUseDescriptionForLogs(t *testing.T) {
	assert := require.New(t)

	read, cancel := newReadFunc(t)
	defer cancel()

	collector := NewErrorCollector("multi error").EnableDescriptionInLogs()
	assert.EqualError(collector, "")

	// skiped
	collector.Append(nil)

	// Logs with description
	collector.Append(New("first"))
	collector.Collect(New("second"), "prefix")

	// Logs without description
	_ = collector.DisableDescriptionInLogs()
	collector.Append(New("third"))

	assert.EqualError(collector, "multi error: first; prefix: second; third")

	checkLogMessages(
		t,
		read(),
		"errors_test.go",
		"multi error: first",
		"multi error: prefix: second",
		"third",
	)
}

func TestAddCallerSkip(t *testing.T) {
	read, cancel := newReadFunc(t)
	defer cancel()

	collector := NewErrorCollector("multi error").AddCallerSkip(-1)
	collector.Append(New("first"))
	collector.Collect(New("second"), "prefix #1")
	collector.Collectf(New("third"), "prefix #2")

	checkLogMessages(
		t,
		read(),
		"errors.go",
		"first",
		"prefix #1: second",
		"prefix #2: third",
	)
}

func TestErrorIs(t *testing.T) {
	assert := require.New(t)
	testErr1 := New("error-1")
	testErr2 := New("error-2")
	testErrPathErr := &fs.PathError{Op: "test", Err: New("fail")}

	collector := NewErrorCollector("multi error")
	collector.Collect(testErr1, "")
	collector.Collect(testErrPathErr, "")
	collector.Collectf(testErr2, "")

	assert.ErrorIs(collector, testErr1)
	assert.ErrorIs(collector, testErr2)
	assert.ErrorIs(collector, testErrPathErr)

	assert.NotErrorIs(collector, New("unknown"))
}

func TestErrorIsNilCollector(t *testing.T) {
	assert := require.New(t)

	var collector *ErrorCollector
	assert.Nil(collector)

	assert.False(Is(collector, io.EOF))
	assert.False(collector.Is(io.EOF))
}

func TestErrorAs(t *testing.T) {
	assert := require.New(t)
	testErrPathErr := &fs.PathError{Op: "test", Err: New("fail")}

	collector := NewErrorCollector("multi error")
	collector.Collect(testErrPathErr, "")

	targetSuccess := &fs.PathError{}
	assert.True(As(collector, &targetSuccess))
	assert.Equal(testErrPathErr, targetSuccess)

	targetFail := new(exec.Error)
	assert.False(As(collector, &targetFail))
	assert.Equal(new(exec.Error), targetFail)
}

func TestErrorAsNilCollector(t *testing.T) {
	assert := require.New(t)

	var collector *ErrorCollector
	assert.Nil(collector)

	target := &fs.PathError{}
	assert.False(As(collector, &target))
	assert.False(collector.As(nil))
}

func TestUnwrap(t *testing.T) {
	assert := require.New(t)
	testErr1 := New("error-1")
	testErr2 := New("error-2")

	collector := NewErrorCollector("multi error")
	assert.Nil(Unwrap(collector))

	collector.Collect(nil, "")
	collector.Collectf(nil, "")

	collector.Collect(testErr1, "")
	assert.Equal(testErr1, Unwrap(collector))
	assert.Equal([]error{testErr1}, Errors(Unwrap(collector)))
	assert.Equal([]error{testErr1}, Errors(collector.GetError()))

	collector.Collect(testErr2, "")
	assert.Equal([]error{testErr1, testErr2}, Errors(Unwrap(collector)))
	assert.Equal([]error{testErr1, testErr2}, Errors(collector.GetError()))
}

func TestUnwrapNilCollector(t *testing.T) {
	assert := require.New(t)

	var collector *ErrorCollector
	assert.Nil(collector)
	assert.Nil(Unwrap(collector))
}

func TestGetErrorNilCollector(t *testing.T) {
	assert := require.New(t)

	var collector *ErrorCollector
	assert.Nil(collector)
	assert.Nil(collector.GetError())
}

func TestGoString(t *testing.T) {
	assert := require.New(t)

	var collector *ErrorCollector
	assert.Nil(collector)
	assert.Equal("(&ErrorCollector)(nil)", fmt.Sprintf("%#v", collector))

	collector = NewErrorCollector("multi error")
	collector.Append(io.EOF)
	assert.Equal(
		`&ErrorCollector{description: multi error, err: [&errors.errorString{s:"EOF"}]}`,
		fmt.Sprintf("%#v", collector),
	)
}

func checkLogMessages(t *testing.T, stdout, callerName string, requiredText ...string) {
	assert := require.New(t)

	lines := strings.Split(stdout, "\n")
	assert.Len(lines, len(requiredText)+1)
	assert.Empty(lines[len(lines)-1])

	for i, text := range requiredText {
		assert.Regexpf(
			fmt.Sprintf(`^I\d+ \d+:\d+:\d+\.\d+\s+\d+ %s:\d+] %s\z`, callerName, text),
			lines[i],
			"line #%d: %q", i, lines[i],
		)
	}
}

func newReadFunc(t *testing.T) (read func() string, cancel func()) {
	assert := require.New(t)

	r, w, err := os.Pipe()
	assert.NoError(err)

	sourceStdOut := os.Stdout
	sourceStdErr := os.Stderr

	os.Stdout = w
	os.Stderr = w

	cancel = func() {
		err := w.Close()
		assert.Truef(err == nil || Is(err, os.ErrClosed), "error: %#v", err)

		os.Stdout = sourceStdOut
		os.Stderr = sourceStdErr
	}

	out := make(chan string, 1)

	go func() {
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, r)
		assert.NoError(err)
		out <- buf.String()
	}()

	read = func() string {
		cancel()
		return <-out
	}

	return read, cancel
}
