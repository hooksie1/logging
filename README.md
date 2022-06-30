# Logging

This is a simple logging package for Go.

## Example

``` go
package main

import (
	"logging"
)

func main() {
	// set level with environment variables
	log := logging.NewLogger()

	// you can also set the level directly
	log.Level = logging.DebugLevel

	log.Info("I'm an info log")
	log.Debug("I'm a debug log")

    // log with format specifiers
	log.Infof("hey there %s", "Gopher")


	log.Error("error always prints")

	// You can also call without defining a logger but this relies on env vars
	// for example this won't print unless LOG_LEVEL=debug is set
	logging.Debugf("hey there again %s", "Gopher")


	log.Fatal("fatal passes through to log.Fatal in std lib")
}
```