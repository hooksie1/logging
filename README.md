# Logging

This is a simple logging package for Go.

## Example

``` go
package main

import (
	"github.com/hooksie1/logging"
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

## Output

Since this uses the default Go logger under the hood, setting an output file instead of stdout is easy.

``` go
package main

import (
	stdlog "log"

	"github.com/hooksie1/logging"
)

func main() {
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		stdlog.Fatal("error")
	}
	defer f.Close()


	log := logging.NewLogger()
	// use SetOutput from std logger
	log.Logger.SetOutput(f)
}
```