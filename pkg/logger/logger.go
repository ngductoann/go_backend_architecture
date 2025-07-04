package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/ngductoann/go_backend_architecture/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerString) *LoggerZap {
	logLevel := config.Log_level
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.FatalLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name, // "/var/log/myapp/foo.log"
		MaxSize:    config.Max_size,      // megabytes
		MaxBackups: config.Max_backups,   // number of backups
		MaxAge:     config.Max_age,       // days
		Compress:   config.Compress,      // compress the backups, disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zapcore.InfoLevel)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
