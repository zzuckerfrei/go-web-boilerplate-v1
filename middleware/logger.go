// fileLogger initializes a zap.Logger that writes to both the console and a specified file.
package middleware

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getConfig() {

}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}
	return encoderConfig

}

func getRotateLogger(filename string) *lumberjack.Logger {
	// Set up lumberjack as a logger:
	rotateLogger := &lumberjack.Logger{
		// Filename:   "./myapp.log", // Or any other path
		Filename:   filename, // Or any other path
		MaxSize:    500,      // MB; after this size, a new log file is created
		MaxBackups: 3,        // Number of backups to keep
		MaxAge:     28,       // Days
		Compress:   true,     // Compress the backups using gzip
	}
	return rotateLogger
}

func GetLogger(filename string) (*zap.Logger, error) {
	// Configure the format
	encoderConfig := getEncoderConfig()

	// Create file and console encoders
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Open the log file
	rotateLogger := getRotateLogger(filename)

	// Create writers for file and console
	writeSyncer := zapcore.AddSync(rotateLogger)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// Set the log level
	// todo setLogLevel
	defaultLogLevel := zapcore.DebugLevel

	// Create cores for writing to the file and console
	fileCore := zapcore.NewCore(fileEncoder, writeSyncer, defaultLogLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, defaultLogLevel)

	// Combine cores
	core := zapcore.NewTee(fileCore, consoleCore)

	// Create the logger with additional context information (caller, stack trace)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	return logger, nil
}
