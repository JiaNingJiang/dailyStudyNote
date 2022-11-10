package main

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestZapLog(t *testing.T) {
	data := &Options{
		LogFileDir: "../log",
		AppName:    "logtool",
		MaxSize:    30,
		MaxBackups: 7,
		MaxAge:     7,
		Config:     zap.Config{},
	}
	data.Development = true
	initLogger(data)
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		logger.Debug(fmt.Sprint("debug log ", i))
		logger.Info(fmt.Sprint("Info log ", i))
		logger.Warn(fmt.Sprint("warn log ", i))
		logger.Error(fmt.Sprint("err log ", i))
	}
}
