package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level 日志等级
type Level string

const (
	// DebugLevel debug等级
	DebugLevel Level = "debug"
	// InfoLevel  info等级
	InfoLevel Level = "info"
	// WarnLevel  warn等级
	WarnLevel Level = "warn"
	// ErrorLevel error等级
	ErrorLevel Level = "error"
	// FatalLevel fatal等级
	FatalLevel Level = "fatal"
	// PanicLevel panic等级
	PanicLevel Level = "panic"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	// EnableColor 是否开启颜色支持
	EnableColor = true
)

func Sync() {
	if logger != nil {
		logger.Sync()
	}
	if sugar != nil {
		sugar.Sync()
	}
}

// Info info级别日志
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Error error级别日志
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Debug debug级别日志
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Warn warn level log
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Fatal  fatal level
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Debugf use fmt.Sprintf to log a templated message  debug level
func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

// Infof use fmt.Sprintf to log a templated message  info level
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

// Warnf use fmt.Sprintf to log a templated message warn level
func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

// Errorf use fmt.Sprintf to log a templated message
func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then call os.exit
func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}

// InitLog 初始化日志 在使用日志之前必须先调用该方法
func InitLog(level string, filePath string) error {
	zl := getZapLevelByName(level)
	outputPath := []string{"stdout"}
	errOutputPath := []string{"stderr"}
	if filePath != "" {
		outputPath = append(outputPath, filePath)
		errOutputPath = append(errOutputPath, filePath)
		EnableColor = false
	}
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zl)
	config.OutputPaths = outputPath
	config.Encoding = "console"
	config.ErrorOutputPaths = errOutputPath
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		//改进版short caller
		encoder.AppendString("(" + caller.TrimmedPath() + ")")
	}
	config.EncoderConfig.ConsoleSeparator = " "
	if EnableColor {
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}
	config.DisableCaller = true
	var err error
	logger, err = config.Build()
	if err != nil {
		return err
	}
	sugar = logger.Sugar()
	return nil
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	case PanicLevel:
		return zapcore.PanicLevel
	default:
		//默认使用info level
		return zapcore.InfoLevel
	}
}

func getLogLevel(name string) Level {
	str := strings.ToLower(name)
	switch str {
	case "debug", "d":
		return DebugLevel
	case "info", "i":
		return InfoLevel
	case "error", "e":
		return ErrorLevel
	case "fatal", "f":
		return FatalLevel
	case "panic", "p":
		return PanicLevel
	default:
		//默认开启info级别日志
		return InfoLevel
	}
}

func getZapLevelByName(level string) zapcore.Level {
	l := getLogLevel(level)
	return getZapLevel(l)
}
