package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewPlanetScaleConfigDefault creates an opinionated zap.Config while
// detecting encoding and Level.
func NewPlanetScaleConfigDefault() Config {
	return NewPlanetScaleConfig(DetectEncoding(), DetectLevel())
}

// NewPlanetScaleConfig creates a zap.Config with the desired encoding and Level.
func NewPlanetScaleConfig(encoding string, level Level) Config {
	return Config{
		Level:    zap.NewAtomicLevelAt(level),
		Encoding: encoding,
		// only buffer the JSON encoder
		Buffered: encoding == JSONEncoding,
	}
}

var defaultEncoderConfig = zapcore.EncoderConfig{
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
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// Config is our logging configration
type Config struct {
	Level    zap.AtomicLevel
	Encoding string
	Buffered bool
}

// Build creates a Logger out of our Config.
// Note that this returns an error, but this doesn't actually
// error. This is maintained for compatibility with
// zapcore.Config{}.Build().
func (cfg Config) Build(opts ...zap.Option) (*Logger, error) {
	var ws zapcore.WriteSyncer = os.Stderr
	// XXX: the internal BufferedWriteSyncer in theory
	// leaks a goroutine for the ticker to flush to stderr,
	// but in practice, this shouldn't particularly be a concern
	// since there we shouldn't be needing to create and destroy
	// loggers at runtime. If this becomes an actual issue
	// we might need to expose a way to get this
	// BufferedWriteSyncer so the caller can call Stop() on it.
	if cfg.Buffered {
		ws = &zapcore.BufferedWriteSyncer{WS: ws}
	}
	log := zap.New(
		zapcore.NewCore(
			cfg.buildEncoder(),
			ws,
			cfg.Level,
		),
		zap.ErrorOutput(ws),
		zap.AddCaller(),
		zap.AddStacktrace(ErrorLevel),
	)
	if len(opts) > 0 {
		log = log.WithOptions(opts...)
	}
	return log, nil
}

func (cfg Config) buildEncoder() zapcore.Encoder {
	// we only suppport pretty or json
	if cfg.Encoding == PrettyEncoding {
		return NewPrettyEncoder(defaultEncoderConfig)
	}
	return zapcore.NewJSONEncoder(defaultEncoderConfig)
}
