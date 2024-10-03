package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// SetupLogging initializes the logger with the appropriate settings.
func SetupLogging() *logrus.Logger {
	// Get the debug from environment variable
	debug := false
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
	if logger == nil {
		logger = logrus.New()
		logger.SetOutput(os.Stdout)
		logger.SetReportCaller(true)

		customFormatter := new(logrus.TextFormatter)
		customFormatter.TimestampFormat = "2006-01-02T15:04:05-0700"
		customFormatter.FullTimestamp = false
		logger.SetFormatter(customFormatter)

		if debug {
			logger.SetLevel(logrus.DebugLevel)
		} else {
			logger.SetLevel(logrus.InfoLevel)
		}
	}

	return logger
}
