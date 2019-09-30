package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func setUpLogger() {
	// perhaps have to set up with fields and organise the logs
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)
	log.Out = os.Stdout
}

func init() {
	setUpLogger()
}

func info(msg string) {
	log.Info(msg)
}

func warn(msg string) {
	log.Warn(msg)
}

func error(msg string) {
	log.Error(msg)
}

func debug(msg string) {
	log.Debug(msg)
}

func fatal(msg string) {
	log.Fatal(msg)
}

func trace(msg string) {
	log.Trace(msg)
}

func panic(msg string) {
	log.Panic(msg)
}

// Message : Log any kind of message - info, warn, error, debug, fatal, trace and panic. It's got levels to it
func Message(msg string, msgType string) {
	switch msgType {
	case "info":
		info(msg)
	case "warn":
		warn(msg)
	case "error":
		error(msg)
	case "debug":
		debug(msg)
	case "fatal":
		fatal(msg)
	case "trace":
		trace(msg)
	case "panic":
		panic(msg)
	default:
		info(msg)
	}
}
