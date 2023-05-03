// Package logger contains logging tools.
package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const permission = 0o600

// InitZap provides logging with zap.
func InitZap(debugMode bool) error {
	logLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	var (
		file *os.File
		err  error
	)

	if debugMode {
		file = os.Stdout
	} else {
		file, err = setFile()
		if err != nil {
			return err
		}
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(file),
		logLevel,
	)

	logger := zap.New(zapcore.NewTee(
		core,
	))

	zap.ReplaceGlobals(logger)

	return nil
}

// config returns EncoderConfig for production environments.
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

// setFile return the location where the log file will be placed.
func setFile() (*os.File, error) {
	dirPath := "."
	fileName := "log.json"
	content := filepath.Join(dirPath, fileName)

	if _, err := os.Stat(content); err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Create(content); err != nil {
				return nil, fmt.Errorf("failed to create file for log output: %w", err)
			}
		} else {
			return nil, fmt.Errorf("a file in a bad state for reasons other than file not existing: %w", err)
		}
	}

	f, err := os.OpenFile(content, os.O_APPEND|os.O_WRONLY, permission)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for log output: %w", err)
	}

	return f, nil
}

// LogDebug is Key-value format debug log.
func LogDebug(msg string, kv ...interface{}) {
	zap.S().Debugw(msg, kv...)
}

// LogErr is Key-value format error log.
func LogErr(msg string, kv ...interface{}) {
	zap.S().Errorw(msg, kv...)
}

// SetLevel sets the log level by specifying a string which
// can be any of:
// ["DEBUG", "INFO", "WARNING", "ERROR", "PANIC", "FATAL"],
// case-insensitive.
func SetLevel(level string, logLevel *zap.AtomicLevel) error {
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
func GetLevel(logLevel *zap.AtomicLevel) zapcore.Level {
	return logLevel.Level()
}
