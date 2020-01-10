// Package golog provides functions for logging in go programs with different
// severity levels and a optional debug mode
package golog

// Libraries
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Log level descriptors
var (
	ok      *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	debug   *log.Logger
)

// Init create the loggers for each log level
func Init(okHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, debugHandle io.Writer, header *string, debugMode bool) {
  var logHeader string
  if header == nil {
    logHeader = ""
  } else {
    logHeader = *header
  }
	ok = log.New(okHandle, logHeader, log.Ldate|log.Ltime)
	info = log.New(infoHandle, logHeader, log.Ldate|log.Ltime)
	warning = log.New(warningHandle, logHeader, log.Ldate|log.Ltime)
	error = log.New(errorHandle, logHeader, log.Ldate|log.Ltime)

	if debugMode {
		debug = log.New(debugHandle, logHeader, log.Ldate|log.Ltime)
		Warn("Enabled Debug logger Mode")
	} else {
		debug = log.New(ioutil.Discard, logHeader, log.Ldate|log.Ltime)
		Debug("Disabled Debug logger Mode")
	}
}

// formatMsgString returns a list of arguments in a string
func formatMsgString(msg ...interface{}) string {
	return fmt.Sprintf(msg[0].(string), msg[1:]...)
}

// Ok prints a OK level message
func Ok(msg ...interface{}) {
	ok.Printf("- OK - %s", formatMsgString(msg...))
}

// Info prints a Info level message
func Info(msg ...interface{}) {
	info.Printf("- INFO - %s", formatMsgString(msg...))
}

// Warn prints a Warning level message
func Warn(msg ...interface{}) {
	warning.Printf("- WARN - %s", formatMsgString(msg...))
}

// Err prints a Error level message
func Err(msg ...interface{}) {
	_, fileName, fileLine, _ := runtime.Caller(1)
	error.Printf("- ERROR - %s:%d:  %s", path.Base(fileName), fileLine, formatMsgString(msg...))
}

// Debug prints a Debug level message
func Debug(msg ...interface{}) {
	_, fileName, fileLine, _ := runtime.Caller(1)
	debug.Printf("- DEBUG - %s:%d:  %s", path.Base(fileName), fileLine, formatMsgString(msg...))
}
