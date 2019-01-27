package main

import (
	"time"

	"go.uber.org/zap/zapcore"
	"os"
	"go.uber.org/zap"
)



func main()  {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("info")

}