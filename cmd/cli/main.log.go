package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1. sugar and logger
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "ngductoann", 25) // like fmt.Printf, but with structured logging
	//
	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "ngductoann"), zap.Int("age", 40)) // provides more structured logging key-value pairs

	// 2. Example of creating different types of loggers
	// logger := zap.NewExample()
	// logger.Info("Hello")
	//
	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")
	//
	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	// 3. Custom logger with configuration
	encoder := getEncoderLog() // format log
	sync := getWriterSync()    // write to file and console
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Hello Custom Logger", zap.String("name", "ngductoann"), zap.Int("age", 25))    // structured logging with custom logger
	logger.Error("Error occurred", zap.String("error", "file not found"), zap.Int("code", 404)) // structured logging with error
}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Set time format to ISO8601 (timestamp format to ISO8601)
	encoderConfig.TimeKey = "time"                        // ts -> time

	// from info INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:24"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig) // or zapcore.NewConsoleEncoder(encoderConfig) for console output
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE, os.ModePerm)

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)                 // sync to stderr
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole) // write to both file and console
}
