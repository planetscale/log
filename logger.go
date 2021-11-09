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

// New creates a new default PlanetScale Logger.
var New = NewPlanetScaleLogger

// NewPlanetScaleSugarLogger creates an opinionated zap.SugaredLogger. Additional customization
// is available by passing in one or more zap.Options.
// NOTE: A SugaredLogger can be converted into a zap.Logger with the .DeSugar() method.
func NewPlanetScaleSugarLogger() *zap.SugaredLogger {
	return NewPlanetScaleLogger().Sugar()
}

// NewPlanetScaleConfig creates an opinionated zap.Config
func NewPlanetScaleConfig() zap.Config {
	encoding := "json"
	if os.Getenv("PS_DEV_MODE") != "" {
		encoding = "console"
	}

	// The default, empty string, unmarshals into "info"
	var level zapcore.Level
	if err := (&level).UnmarshalText([]byte(os.Getenv("PS_LOG_LEVEL"))); err != nil {
		panic("Invalid PS_LOG_LEVEL value: " + os.Getenv("PS_LOG_LEVEL"))
	}

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
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

type Logger = zap.Logger

// Re-export of all zap field functions for convenience
var (
	Any         = zap.Any
	Array       = zap.Array
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	Bools       = zap.Bools
	ByteString  = zap.ByteString
	ByteStrings = zap.ByteStrings
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex128s = zap.Complex128s
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Complex64s  = zap.Complex64s
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Durations   = zap.Durations
	Error       = zap.Error
	Errors      = zap.Errors
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Float32s    = zap.Float32s
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float64s    = zap.Float64s
	Inline      = zap.Inline
	Int         = zap.Int
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int16s      = zap.Int16s
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int32s      = zap.Int32s
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int64s      = zap.Int64s
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	Int8s       = zap.Int8s
	Intp        = zap.Intp
	Ints        = zap.Ints
	NamedError  = zap.NamedError
	Namespace   = zap.Namespace
	Object      = zap.Object
	Reflect     = zap.Reflect
	Skip        = zap.Skip
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	String      = zap.String
	Stringer    = zap.Stringer
	Stringp     = zap.Stringp
	Strings     = zap.Strings
	Time        = zap.Time
	Timep       = zap.Timep
	Times       = zap.Times
	Uint        = zap.Uint
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint16s     = zap.Uint16s
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint32s     = zap.Uint32s
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint64s     = zap.Uint64s
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uint8s      = zap.Uint8s
	Uintp       = zap.Uintp
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Uintptrs    = zap.Uintptrs
	Uints       = zap.Uints
)
