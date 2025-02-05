package core

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.WithField("app", "redis-mini")
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetOutput(os.Stdout)
}
