package main

import (
	"github.com/learn-go/Exercise/day2/example/example1_2/loggrt"
	"go.uber.org/zap"
)

func main() {
	loggrt.Init()
	loggrt.Log.Debug("test",zap.String("key","value"))
}