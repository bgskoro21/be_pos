package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetupLogger(){
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
