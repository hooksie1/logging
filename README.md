# Logging

This is a simple logging package for Go.

## Example

``` go
package main

import (
	"logging"
	"time"
)

func main() {
    log := logging.NewLogger()

    // level can be set as an environment varirable or through setting the level manually
    log.Level = logging.DebugLevel

    log.Info("I'm an info log")
    time.Sleep(1 * time.Second)
    log.Debug("I'm a debug log")

    // You can also call the logger without defining a logger
    logging.Infof("hey there %s", "Gopher")
}
```