package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap.SugaredLogger
}

func NewLogger(file, lvl string) Logger {
	cfgConsole := zap.NewProductionEncoderConfig()
	cfgConsole.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfgConsole.ConsoleSeparator = " "
	cfgConsole.EncodeTime = zapcore.TimeEncoderOfLayout("060102150405")
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0750)
	if err != nil {
		panic(err)
	}
	al, err := zap.ParseAtomicLevel(lvl)
	if err != nil {
		panic(err)
	}
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(cfgConsole),
			zapcore.Lock(f),
			al.Level(),
		),
	)
	zl := zap.New(core, zap.AddCaller())
	defer zl.Sync()
	zs := zl.Sugar()
	l := Logger{*zs}
	return l
}
