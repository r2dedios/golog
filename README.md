# GoLog
Golang Logging custom module

![Release](https://img.shields.io/github/v/release/r2dedios/golog?color=purple&include_prereleases&label=Release)
[![Go Report Card](https://goreportcard.com/badge/github.com/r2dedios/golog)](https://goreportcard.com/report/github.com/r2dedios/golog)

## Package
Package golog provides functions for logging in go programs with different
severity levels and a optional debug mode
```go
import "github.com/r2dedios/golog"
```


## Functions

```
func Debug(msg ...interface{})
    Debug prints a Debug level message

func Err(msg ...interface{})
    Err prints a Error level message

func Info(msg ...interface{})
    Info prints a Info level message

func Init(okHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, debugHandle io.Writer, header *string, debugMode bool)
    Init create the loggers for each log level

func Ok(msg ...interface{})
    Ok prints a OK level message

func Warn(msg ...interface{})
    Warn prints a Warning level message
```
