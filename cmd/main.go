package main

import (
	"logging"
	"time"
)

func main() {
	// set level with env var
	log := logging.NewLogger()

	log.Info("I'm an info log")
	time.Sleep(1 * time.Second)
	log.Debug("I'm a debug log")
	log.Infof("hey there %s", "Gopher")

	// You can also call without defining a logger
	logging.Debugf("hey there again %s", "Gopher")
}
