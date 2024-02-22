package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

var logLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

func InitLogger(logLevel string) {
	level, ok := logLevelMap[logLevel]
	if !ok {
		fmt.Printf("Invalid log level: %s. Setting default to 'info'\n", logLevel)
		level = logrus.InfoLevel
	}

	Logger.SetLevel(level)
	Logger.SetFormatter(&logrus.JSONFormatter{})
}
