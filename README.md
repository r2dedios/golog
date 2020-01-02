# GoLog
Golang Logging custom module

![Release](https://img.shields.io/github/v/release/r2dedios/golog?color=purple&include_prereleases&label=Release)
[![Go Report Card](https://goreportcard.com/badge/github.com/r2dedios/golog)](https://goreportcard.com/report/github.com/r2dedios/golog)

## Package
Package golog provides functions for logging in go programs with different
severity levels and a optional debug mode
```go
import "golog"
```


## Functions

```go
// DebugMsg prints a Debug level message
func DebugMsg(msg ...interface{})

// ErrMsg prints a Error level message
func ErrMsg(msg ...interface{})

// InfoMsg prints a Info level message
func InfoMsg(msg ...interface{})

// Init create the loggers for each log level
func Init(okHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, debugHandle io.Writer, header string, debugMode bool)

// OkMsg prints a OK level message
func OkMsg(msg ...interface{})

// WarnMsg prints a Warning level message
func WarnMsg(msg ...interface{})

```
