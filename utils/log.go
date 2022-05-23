package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Save file log cut
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./migration.log", // Log name
		MaxSize:    10,                // File content size, MB
		MaxBackups: 5,                 // Maximum number of old files retained
		MaxAge:     30,                // Maximum number of days to keep old files
		Compress:   false,             // Is the file compressed
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// The format time can be customized
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitializeLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		zapcore.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller())
}
