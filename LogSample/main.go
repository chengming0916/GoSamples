package main

import (
	_ "GoSamples/LogSample/pkg/log" // 引用log,使配置生效
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Debugln("Debug Sample")
	logrus.Infoln("Info Sample")
	logrus.Warnln("Warn Sample")
	logrus.Errorln("Error Sample")
	// logrus.Fatalln("Fatal Sample") // fatal会导致程序退出

	// 结构化日志
	logrus.WithFields(logrus.Fields{
		"field-1": "field-1 content",
		"field-2": 1,
		"field-3": time.Now().Local().String(),
	}).Infoln("Struct Sample")
}
