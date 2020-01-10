package golog

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

var dateRegex = `([12]\d{3}/(0[1-9]|1[0-2])/(0[1-9]|[12]\d|3[01])) (0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]`

// TestInit checks that Init function assings the correct descriptors for each level
func TestInit(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	Init(stdout, stdout, stdout, stderr, stdout, nil, true)
	if ok.Writer() != stdout {
		t.Error("OK level output is not correct")
	}

	if info.Writer() != stdout {
		t.Error("Info level output is not correct")
	}

	if warning.Writer() != stdout {
		t.Error("Warning level output is not correct")
	}

	if error.Writer() != stderr {
		t.Error("Error level output is not correct")
	}

	if debug.Writer() != stdout {
		t.Error("Debug level output is not correct")
	}
}

// TestOk checks the OK level message function
func TestOk(t *testing.T) {
	buffer := &bytes.Buffer{}
	ok.SetOutput(buffer)
	entry := "ok level message"
	Ok(entry)
	testRegex := fmt.Sprintf("^%s - OK - %s\\n$", dateRegex, entry)
	matched, _ := regexp.MatchString(testRegex, buffer.String())
	if !matched {
		t.Error("Ok level message function fails")
	}
}

// TestInfo checks the Informative level message function
func TestInfo(t *testing.T) {
	buffer := &bytes.Buffer{}
	info.SetOutput(buffer)
	entry := "info level message"
	Info(entry)
	testRegex := fmt.Sprintf("^%s - INFO - %s\\n$", dateRegex, entry)
	matched, _ := regexp.MatchString(testRegex, buffer.String())
	if !matched {
		t.Error("Info level message function fails")
	}
}

// TestWarn checks the Warning level message function
func TestWarn(t *testing.T) {
	buffer := &bytes.Buffer{}
	warning.SetOutput(buffer)
	entry := "warning level message"
	Warn(entry)
	testRegex := fmt.Sprintf("^%s - WARN - %s\\n$", dateRegex, entry)
	matched, _ := regexp.MatchString(testRegex, buffer.String())
	if !matched {
		t.Error("Warning level message function fails")
	}
}

// TestErr checks the Error level message function
func TestErr(t *testing.T) {
	buffer := &bytes.Buffer{}
	error.SetOutput(buffer)
	entry := "error level message"
	Err(entry)
	testRegex := fmt.Sprintf("^%s - ERROR - \\w*.go:\\d*:  %s\\n$", dateRegex, entry)
	matched, _ := regexp.MatchString(testRegex, buffer.String())
	if !matched {
		t.Error("Error level message function fails")
	}
}

// TestDebug checks the Debug level message function
func TestDebug(t *testing.T) {
	buffer := &bytes.Buffer{}
	debug.SetOutput(buffer)
	entry := "debug level message"
	Debug(entry)
	testRegex := fmt.Sprintf("^%s - DEBUG - \\w*.go:\\d*:  %s\\n$", dateRegex, entry)
	matched, _ := regexp.MatchString(testRegex, buffer.String())
	if !matched {
		t.Error("Debug level message function fails")
	}
}
