package log

import (
	"io"
	"os"
	"runtime"
	"time"

	"log/slog"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"k8s.io/klog/v2"
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

	writer1 := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    500, // 日志文件大小,单位MB
		MaxBackups: 3,   // 最多保留日志文件数
		MaxAge:     30,  // 最大保存天数
	}

	// v1.21版本官方日志库
	handler := slog.NewTextHandler(writer1, &slog.HandlerOptions{Level: slog.LevelDebug})
	// handler := slog.NewJSONHandler(writer1, &slog.HandlerOptions{Level: slog.LevelDebug}) // Json
	slog.SetDefault(slog.New(handler))

	// klog
	klog.InitFlags(nil) // 初始化全局标志
	klog.SetOutput(io.MultiWriter(os.Stdout, writer1))

	// zap.NewProduction() // 输出Json日志，调试级别不输出
	// zap.NewDevelopment() // 输出行日志
	// log := zap.NewExample() // 输出Json日志
	// log.Info("", zap.Int("age", 4), zap.String("gender", "M")) // 记录结构化日志
	// 以printf格式记录语句
	// log.Sugar().Debug() //
	// log.Sugar().Desugar() // 从sugar logger切换到标准记录器

	// encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	// encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写标记日志级别
	// encoder := zapcore.NewConsoleEncoder(encoderConfig)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer1)), // 输出到控制台和文件
		level, // 日志级别
	)

	logger := zap.New(core, zap.AddCaller()) // AddCaller 将调用函数信息记录到日志中（堆栈跟踪）
	// sugarLogger := logger.Sugar()
	defer logger.Sync() // 确保日志缓冲区被刷新
}
