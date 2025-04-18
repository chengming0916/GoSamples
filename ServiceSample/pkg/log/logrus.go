package log

import (
	"io"
	"os"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	FieldMap: logrus.FieldMap{
	// 		logrus.FieldKeyMsg:   "message",
	// 		logrus.FieldKeyLevel: "level",
	// 		logrus.FieldKeyTime:  "time",
	// 	},
	// 	TimestampFormat: time.RFC3339,
	// })

	logrus.SetFormatter(&logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339,
	})

	var filePath string
	if runtime.GOOS == "linux" {
		filePath = "/var/log/go-service-sample.log"
	} else if runtime.GOOS == "windows" {
		filePath = "./logs/go-service-sample.log"
	}

	writer, _ := rotatelogs.New(
		filePath+".%Y%m%d",
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithRotationSize(50*1024*1024),
	)
	logrus.SetOutput(io.MultiWriter(os.Stdout, writer))
}
