package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

var logger = createLogger()

type LogWriter struct {
	*logrus.Logger
}

func (l *LogWriter) Write(p []byte) (int, error) {
	l.Error(string(p))
	return len(p), nil
}

func createLogger() *logrus.Logger {
	res := logrus.New()

	logWriter := LogWriter{res}
	log.SetFlags(0)
	log.SetOutput(&logWriter)

	return res
}

func initLogger() {
	if viper.GetBool("production") {
		logger.Formatter = new(logrus.JSONFormatter)
	} else {
		logger.Level = logrus.DebugLevel
	}
}
