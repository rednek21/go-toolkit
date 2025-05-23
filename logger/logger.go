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

	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
	PANIC = "panic"
)

func New(cfg *Config) (*zap.Logger, error) {
	level, err := parseLogLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	encoder := getEncoder(cfg.Format)

	var cores []zapcore.Core

	if cfg.FilePath != "" {
		fileSyncer := zapcore.Lock(zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.MaxSizeMB,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAgeDays,
			Compress:   cfg.Compress,
		}))
		cores = append(cores, zapcore.NewCore(encoder, fileSyncer, level))
	}

	if cfg.Output == STDOUT || cfg.Output == STDERR {
		consoleSyncer := zapcore.Lock(os.Stdout)
		if cfg.Output == STDERR {
			consoleSyncer = zapcore.Lock(os.Stderr)
		}
		cores = append(cores, zapcore.NewCore(encoder, consoleSyncer, level))
	}

	if len(cores) == 0 {
		cores = append(cores, zapcore.NewCore(getEncoder(JSON), zapcore.AddSync(os.Stdout), zapcore.InfoLevel))
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func parseLogLevel(levelStr string) (zapcore.Level, error) {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(levelStr))
	return level, err
}

func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if format == JSON {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
