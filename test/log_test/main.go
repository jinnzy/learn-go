package main

import (
	"go.uber.org/zap"
	"time"
	"go.uber.org/zap/zapcore"
		"os"
)


var log *zap.SugaredLogger
// 时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
// zap日志配置信息
func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func qnit() () {
	//loglevel := setting.ServerSetting.RunMode
	// 读取配置文件
	// 初始化结构体，解析配置文件

	// 根据配置文件赋值
	loglevel := "debug"
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(newEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		level,
	)
	// 要改json格式的话，把NewConsoleEncoder 改成NewJSONEncoder
	// 改成文件输出的话修改zapcore.AddSync(os.Stdout))
	// 连接 https://blog.csdn.net/NUCEMLS/article/details/86534444
	logger := zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1)) // AddCallerSkip(1) 跳过封装函数的调用
	logger.Info("logger",zap.String("key","value"))
}


func main()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	qnit()
}