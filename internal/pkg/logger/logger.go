// Package logger contains logging tools.
package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevel zap.AtomicLevel

// InitZap provides logging with zap.
func InitZap() {
	logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// Standard output
	stdCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	logger := zap.New(zapcore.NewTee(
		stdCore,
	))

	zap.ReplaceGlobals(logger)
}

// config returns EncoderConfig for production environments
func config() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.MessageKey = "msg"
	cfg.LevelKey = "level"
	cfg.NameKey = "name"
	cfg.TimeKey = "timestamp"
	cfg.CallerKey = "caller"
	cfg.FunctionKey = "func"
	cfg.StacktraceKey = "stacktrace"
	cfg.LineEnding = "\n"
	cfg.EncodeTime = zapcore.EpochTimeEncoder
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder

	return cfg
}

// LogDebug is Key-value format debug log
func LogDebug(msg string, kv ...interface{}) {
	zap.S().Debugw(msg, kv...)
}

// LogErr is Key-value format error log
func LogErr(msg string, kv ...interface{}) {
	zap.S().Errorw(msg, kv...)
}

// SetLevel sets the log level by specifying a string which
// can be any of:
// ["DEBUG", "INFO", "WARNING", "ERROR", "PANIC", "FATAL"],
// case-insensitive.
func SetLevel(level string) error {
	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		logLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		fallthrough
	case "WARNING":
		logLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		logLevel.SetLevel(zapcore.ErrorLevel)
	case "PANIC":
		logLevel.SetLevel(zapcore.PanicLevel)
	case "FATAL":
		logLevel.SetLevel(zapcore.FatalLevel)
	default:
		return fmt.Errorf("invalid log level string: %v", level)
	}

	return nil
}

// GetLevel returns the current log level.
func GetLevel() zapcore.Level {
	return logLevel.Level()
}
