package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewPlanetScaleLogger creates an opinionated zap.Logger. Additional customization
// is available by passing in one or more zap.Options.
func NewPlanetScaleLogger() *zap.Logger {
	logger, err := NewPlanetScaleConfig().Build()
	if err != nil {
		panic("Unexpected error initializing PlanetScale logger: " + err.Error())
	}
	return logger
}

// NewPlanetScaleSugarLogger creates an opinionated zap.SugaredLogger. Additional customization
// is available by passing in one or more zap.Options.
// NOTE: A SugaredLogger can be converted into a zap.Logger with the .DeSugar() method.
func NewPlanetScaleSugarLogger() *zap.SugaredLogger {
	logger, err := NewPlanetScaleConfig().Build()
	if err != nil {
		panic("Unexpected error initializing PlanetScale logger: " + err.Error())
	}
	return logger.Sugar()
}

// NewPlanetScaleConfig creates an opinionated zap.Config
func NewPlanetScaleConfig() zap.Config {
	encoding := "json"
	if os.Getenv("PS_DEV_MODE") != "" {
		encoding = "console"
	}

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
