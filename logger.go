package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewPlanetScaleLogger creates an opinionated zap.Logger. Additional customization
// is available by passing in one or more zap.Options.
func NewPlanetScaleLogger() *Logger {
	return NewPlanetScaleLoggerAtLevel(DetectLevel())
}

// NewPlanetScaleLoggerAtLevel creates an opinionated Logger at a desired Level.
func NewPlanetScaleLoggerAtLevel(l Level) *Logger {
	logger, err := NewPlanetScaleConfig(DetectEncoding(), l).Build()
	if err != nil {
		panic("Unexpected error initializing PlanetScale logger: " + err.Error())
	}
	return logger
}

// New creates a new default PlanetScale Logger with auto detection of level.
func New() *Logger {
	return NewPlanetScaleLogger()
}

// NewAtLevel creates a new PlanetScale Logger at the desired Level.
func NewAtLevel(l Level) *Logger {
	return NewPlanetScaleLoggerAtLevel(l)
}

// NewPlanetScaleSugarLogger creates an opinionated zap.SugaredLogger. Additional customization
// is available by passing in one or more zap.Options.
// NOTE: A SugaredLogger can be converted into a zap.Logger with the .DeSugar() method.
func NewPlanetScaleSugarLogger() *SugaredLogger {
	return NewPlanetScaleLogger().Sugar()
}

// NewNop returns a no-op logger
func NewNop() *Logger {
	return zap.NewNop()
}

// DetectEncoding detects the encoding to use based on PS_DEV_MODE env var.
func DetectEncoding() string {
	if os.Getenv("PS_DEV_MODE") != "" {
		return PrettyEncoding
	}
	return JSONEncoding
}

// DetectBuffering detects both if an override is set with PS_LOG_BUFFERED,
// and what the override says. The common case would be to entirely disable
// buffering, and simply setting PS_LOG_BUFFERED=0 would explicitly disable it.
func DetectBuffering() (bool, bool) {
	v, isSet := os.LookupEnv("PS_LOG_BUFFERED")
	return v == "1", isSet
}

// ParseLevel parses a level based on the lower-case or all-caps ASCII
// representation of the log level. If the provided ASCII representation is
// invalid an error is returned.
//
// This is particularly useful when dealing with text input to configure log
// levels.
//
// This is vendored out of `zapcore` since it's added in newer versions, so
// it's trivial enough to vendor and not require a newer `zap` module.
func ParseLevel(text string) (Level, error) {
	var level Level
	err := level.UnmarshalText([]byte(text))
	return level, err
}

// DetectLevel returns a the Level based on PS_LOG_LEVEL env var.
func DetectLevel() Level {
	// The default, empty string, unmarshals into "info"
	level, err := ParseLevel(os.Getenv("PS_LOG_LEVEL"))
	if err != nil {
		panic("Invalid PS_LOG_LEVEL value: " + os.Getenv("PS_LOG_LEVEL"))
	}
	return level
}

const (
	PrettyEncoding = "pretty"
	JSONEncoding   = "json"
)

type (
	Logger        = zap.Logger
	SugaredLogger = zap.SugaredLogger
	Field         = zap.Field
	Level         = zapcore.Level
)

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)

var WithCaller = zap.WithCaller

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

// Stringers constructs a field with the given key, holding a list of the output provided by the value's String method
func Stringers[T interface{ String() string }](key string, value []T) Field {
	return zap.Stringers(key, value)
}
