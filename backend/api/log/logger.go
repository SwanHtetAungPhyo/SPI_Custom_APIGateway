package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func SetupLogger() {

	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}
func GetLogger() *logrus.Logger {
	return log
}
