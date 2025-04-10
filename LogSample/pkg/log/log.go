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
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "0000-00-00 00:00:00",
	})

	// logger.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors: true, // 强制使用颜色输出
	// 	// DisableColors: true, //禁用颜色输出
	// 	ForceQuote: true, // 强制引用所有值
	// 	// DisableQuote: true, // 禁用引用所有值
	// 	// DisableTimestamp: true, // 禁用时间戳记录
	// 	TimestampFormat: "0000-00-00 00:00:00",
	// 	FullTimestamp:   true, // 连接TTY时输出完整时间戳
	// })

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	var filePath string
	if runtime.GOOS == "linux" {
		filePath = "/var/log/logrus-sample.log"
	} else if runtime.GOOS == "windows" {
		os.Mkdir("./logs", os.FileMode(os.O_CREATE))
		filePath = "./logs/logrus-sample.log"
	}

	// 日志输出到文件
	// file, _ := os.OpenFile("./logs/logrus-sample.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// logrus.SetOutput(io.MultiWriter(file, os.Stdout)) // 同时输出到控制台和文件

	// 日志分割输出
	writer, _ := rotatelogs.New(
		filePath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filePath),                     // 为最新的日志建立软连接
		rotatelogs.WithMaxAge(time.Duration(108)*time.Second), // 设置文件清理前的最长保存时间,WithMaxAge 和 WithRotationCount二者只能设置一个
		//rotatelogs.WithRotationCount(10),                           // 设置文件清理前的最多保存文件数,
		//rotatelogs.WithRotationSize(5*1204),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second), // 设置日志分割的时间，隔多久分割一次
	)
	logrus.SetOutput(io.MultiWriter(os.Stdout, writer))
}
