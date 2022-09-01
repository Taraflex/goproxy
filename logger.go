package goproxy

import (
	"log"
	"os"
)

type Logger interface {
	Printf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
}

type DefaultLogger struct{}

var stdout = log.New(os.Stdout, "INFO: ", log.LstdFlags)
var stderr = log.New(os.Stderr, "WARN: ", log.LstdFlags)

func (*DefaultLogger) Printf(format string, v ...interface{}) {
	stdout.Printf(format, v...)
}

func (*DefaultLogger) Warnf(format string, v ...interface{}) {
	stderr.Printf(format, v...)
}
