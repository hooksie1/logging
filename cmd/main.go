package main

import (
	"fmt"
	"net/http"

	"github.com/hooksie1/logging"
)

func context(r *http.Request) string {
	msg := fmt.Sprintf("Method: %s", r.Method)
	return msg
}

func main() {
	r := http.Request{
		Method: "GET",
	}

	// set level with environment variables
	log := logging.NewLogger()

	// you can also set the level directly
	log.Level = logging.InfoLevel

	log.Info("I'm an info log")
	log.Debug("I'm a debug log")

	// log with format specifiers
	log.Infof("hey there %s", "Gopher")

	log.Error("error always prints")

	// You can also call without defining a logger but this relies on env vars
	// for example this won't print unless LOG_LEVEL=debug is set
	logging.Debugf("hey there again %s", "Gopher")

	// You can use a context to log specific contexts around messages.
	// For example, to log the method of an http.Request, you create the logger with context
	// and pass some string in. Here we pass in a func that parses an http.Request to return the method
	lc := log.WithContext(context(&r))
	lc.Info("hey there")

	// Fatal passes through to the std lib logging Fatal method
	log.Fatal("fatal passes through to log.Fatal in std lib")
}
