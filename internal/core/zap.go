package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func Zap() (sugar *zap.SugaredLogger) {
	fileAll := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "all.log",
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})

	consoleStdout := zapcore.Lock(zapcore.AddSync(os.Stdout))
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileAll, zapcore.DebugLevel),
		zapcore.NewCore(encoder, consoleStdout, zapcore.DebugLevel),
	)
	logger := zap.New(core)
	defer logger.Sync()
	sugar = logger.Sugar()
	return sugar
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
