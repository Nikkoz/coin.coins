package logger

import (
	"coins/configs"
	"coins/pkg/types/context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Logger struct {
	logger *zap.Logger
}

func new(cfg configs.Config) (*Logger, error) {
	var config zap.Config

	if cfg.App.Environment.IsProduction() {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	level := level(cfg.Log)
	config.Level = zap.NewAtomicLevelAt(level)

	newLogger, err := config.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	newLogger.Info("Set LOG_LEVEL", zap.Stringer("level", level))

	log = &Logger{logger: newLogger}

	return log, nil
}

func NewLogger(cfg configs.Config) (*Logger, error) {
	return new(cfg)
}

func level(cfg configs.Log) zapcore.Level {
	switch strings.ToUpper(strings.TrimSpace(string(cfg.Level))) {
	case "ERR", "ERROR":
		return zapcore.ErrorLevel
	case "WARN", "WARNING":
		return zapcore.WarnLevel
	case "INFO":
		return zapcore.InfoLevel
	case "DEBUG":
		return zapcore.DebugLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *Logger) getContextFields(ctx context.Context) []zap.Field {
	return []zap.Field{zap.String("id", ctx.ID())}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) DebugWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) InfoWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) WarnWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Warn(msg, fields...)
}

func (l *Logger) Error(msg interface{}, fields ...zap.Field) {
	if msg == nil {
		return
	}

	switch v := msg.(type) {
	case string:
		l.logger.Error(v, fields...)
	case error:
		l.logger.Error(v.Error(), fields...)
	case fmt.Stringer:
		l.logger.Error(v.String(), fields...)
	default:
		l.logger.Error(fmt.Sprintf("%v", v), fields...)
	}
}

func (l *Logger) ErrorWithContext(ctx context.Context, err error, fields ...zap.Field) error {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Error(err, fields...)

	return err
}

func (l *Logger) Fatal(msg interface{}, fields ...zap.Field) {
	if msg == nil {
		return
	}

	switch msg.(type) {
	case string:
		if v, ok := msg.(string); ok {
			l.logger.Fatal(v, fields...)
		}
	case error:
		if v, ok := msg.(error); ok {
			l.logger.Fatal(v.Error(), fields...)
		}
	case fmt.Stringer:
		if v, ok := msg.(fmt.Stringer); ok {
			l.logger.Fatal(v.String(), fields...)
		}
	default:
		l.logger.Fatal(fmt.Sprintf("%v", msg), fields...)
	}
}

func (l *Logger) FatalWithContext(ctx context.Context, err error, fields ...zap.Field) error {
	fields = append(fields, l.getContextFields(ctx)...)

	l.Fatal(err, fields...)

	return err
}
