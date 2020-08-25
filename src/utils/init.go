package utils

import "github.com/sirupsen/logrus"

func InitLog() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		TimestampFormat: "2006/01/02 15:04:05",
		FullTimestamp:   true,
	})
}
