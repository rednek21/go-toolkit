package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const (
	JSON = "json"
	YAML = "yaml"

	STDOUT = "stdout"
	STDERR = "stderr"
)

//func NewLogger(logFile string, level zapcore.Level) (*zap.Logger, error) {
//	config := zap.NewProductionConfig()
//	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//
//	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
//	fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
//
//	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		return nil, err
//	}
//
//	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
//	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level)
//
//	return zap.New(zapcore.NewTee(consoleCore, fileCore),
//		zap.AddCaller(),
//		zap.AddStacktrace(zap.ErrorLevel),
//	), nil
//}

func New(cfg *Config) (*zap.Logger, error) {
	level, err := parseLogLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	encoder := getEncoder(cfg.Format)

	var cores []zapcore.Core

	if cfg.FilePath != "" {
		fyleSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.MaxSizeMB,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAgeDays,
			Compress:   cfg.Compress,
		})
		cores = append(cores, zapcore.NewCore(encoder, fyleSyncer, level))
	}

	if cfg.Output == STDOUT || cfg.Output == STDERR {
		consoleSyncer := zapcore.Lock(os.Stdout)
		if cfg.Output == STDERR {
			consoleSyncer = zapcore.Lock(os.Stderr)
		}
		cores = append(cores, zapcore.NewCore(encoder, consoleSyncer, level))
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(logger)
	return logger, nil
}

// parseLogLevel приводит строковый level в zapcore.Level
func parseLogLevel(levelStr string) (zapcore.Level, error) {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(levelStr))
	return level, err
}

// getEncoder возвращает кодировщик для JSON или console
func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if format == JSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
