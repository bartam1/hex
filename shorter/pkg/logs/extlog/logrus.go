package extlog

import (
	"strconv"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init() {
	//Log callers name too
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv, _ := strconv.ParseBool("false"); isLocalEnv { //os.Getenv("LOCAL_ENV")
		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}

	logrus.SetLevel(logrus.DebugLevel)
}
