package log

import (
	"github.com/golang/glog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewPlanetScaleLogger creates an opinionated zap.Logger. Additional customization
// is available by passing in one or more zap.Options.
func NewPlanetScaleLogger(opts ...zap.Option) (*zap.Logger, error) {
	return NewPlanetScaleConfig().Build(opts...)
}

// NewPlanetScaleSugarLogger creates an opinionated zap.SugaredLogger. Additional customization
// is available by passing in one or more zap.Options.
// NOTE: A SugaredLogger can be converted into a zap.Logger with the .DeSugar() method.
func NewPlanetScaleSugarLogger(opts ...zap.Option) (*zap.SugaredLogger, error) {
	logger, err := NewPlanetScaleConfig().Build(opts...)
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

// NewPlanetScaleConfig creates an opinionated zap.Config
func NewPlanetScaleConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
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

// TODO doc
func HijackGlog(logger *zap.Logger) {
	slogger := logger.Sugar()
	glog.SetLogger(&glog.LoggerFunc{
		DebugfFunc: func(f string, a ...interface{}) { slogger.Debugf(f, a...) },
		InfofFunc:  func(f string, a ...interface{}) { slogger.Infof(f, a...) },
		WarnfFunc:  func(f string, a ...interface{}) { slogger.Warnf(f, a...) },
		ErrorfFunc: func(f string, a ...interface{}) { slogger.Errorf(f, a...) },
	})
}
