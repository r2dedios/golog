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
func Init(okHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, debugHandle io.Writer, header string, debugMode bool) {
	ok = log.New(okHandle, header, log.Ldate|log.Ltime)
	info = log.New(infoHandle, header, log.Ldate|log.Ltime)
	warning = log.New(warningHandle, header, log.Ldate|log.Ltime)
	error = log.New(errorHandle, header, log.Ldate|log.Ltime)

	if debugMode {
		debug = log.New(debugHandle, header, log.Ldate|log.Ltime)
		WarnMsg("Enabled Debug logger Mode")
	} else {
		debug = log.New(ioutil.Discard, header, log.Ldate|log.Ltime)
		DebugMsg("Disabled Debug logger Mode")
	}
}

// formatMsgString returns a list of arguments in a string
func formatMsgString(msg ...interface{}) string {
	return fmt.Sprintf(msg[0].(string), msg[1:]...)
}

// OkMsg prints a OK level message
func OkMsg(msg ...interface{}) {
	ok.Printf("- OK - %s", formatMsgString(msg...))
}

// InfoMsg prints a Info level message
func InfoMsg(msg ...interface{}) {
	info.Printf("- INFO - %s", formatMsgString(msg...))
}

// WarnMsg prints a Warning level message
func WarnMsg(msg ...interface{}) {
	warning.Printf("- WARN - %s", formatMsgString(msg...))
}

// ErrMsg prints a Error level message
func ErrMsg(msg ...interface{}) {
	_, fileName, fileLine, _ := runtime.Caller(1)
	error.Printf("- ERROR - %s:%d:  %s", path.Base(fileName), fileLine, formatMsgString(msg...))
}

// DebugMsg prints a Debug level message
func DebugMsg(msg ...interface{}) {
	_, fileName, fileLine, _ := runtime.Caller(1)
	debug.Printf("- DEBUG - %s:%d:  %s", path.Base(fileName), fileLine, formatMsgString(msg...))
}
