//nolint:errcheck
package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var bufferpool = buffer.NewPool()

var _encPool = sync.Pool{New: func() interface{} {
	return &prettyEncoder{}
}}

func getEncoder() *prettyEncoder {
	return _encPool.Get().(*prettyEncoder)
}

func putEncoder(enc *prettyEncoder) {
	enc.buf = nil
	_encPool.Put(enc)
}

type prettyEncoder struct {
	start time.Time
	buf   *buffer.Buffer
}

func NewPrettyEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return &prettyEncoder{
		start: time.Now(),
		buf:   bufferpool.Get(),
	}
}

func (enc *prettyEncoder) clone() *prettyEncoder {
	clone := getEncoder()
	clone.start = enc.start
	clone.buf = bufferpool.Get()
	return clone
}

func (enc *prettyEncoder) Clone() zapcore.Encoder {
	clone := enc.clone()
	clone.buf.Write(enc.buf.Bytes())
	return clone
}

func (enc *prettyEncoder) addAttribute(as ...attribute) {
	enc.buf.AppendString(escape)
	enc.buf.AppendByte('[')
	first := true
	for _, a := range as {
		if first {
			first = false
		} else {
			enc.buf.AppendByte(';')
		}
		enc.buf.AppendInt(int64(a))
	}
	enc.buf.AppendByte('m')
}

func (enc *prettyEncoder) addKey(key string) {
	enc.buf.AppendByte(' ')
	enc.addAttribute(attributeFgGreen)
	enc.buf.AppendString(key)
	enc.addAttribute(attributeReset)
	enc.buf.AppendByte('=')
}

func (enc *prettyEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	final := enc.clone()

	final.addAttribute(attributeFgWhite)
	fmt.Fprintf(final.buf, "%13s", time.Since(final.start))
	final.addAttribute(attributeReset)

	final.buf.AppendString(" |")

	switch ent.Level {
	case DebugLevel:
		final.addAttribute(attributeFgMagenta)
	case InfoLevel:
		final.addAttribute(attributeFgCyan)
	case WarnLevel:
		final.addAttribute(attributeFgYellow)
	case ErrorLevel:
		final.addAttribute(attributeFgRed)
	case DPanicLevel, PanicLevel:
		final.addAttribute(attributeBgRed)
	case FatalLevel:
		final.addAttribute(attributeBgHiRed, attributeFgHiWhite)
	}
	final.buf.AppendString(strings.ToUpper(ent.Level.String())[:4])
	final.addAttribute(attributeReset)
	final.buf.AppendByte('|')

	final.buf.AppendByte(' ')
	final.buf.AppendString(ent.Message)

	if ent.LoggerName != "" {
		final.AddString("logger", ent.LoggerName)
	}

	if ent.Caller.Defined {
		final.addKey("caller")
		zapcore.ShortCallerEncoder(ent.Caller, final)
	}

	// Write prior fields accumulated from `With()` calls first before writing our new fields.
	final.buf.Write(enc.buf.Bytes())

	addFields(final, fields)

	if ent.Stack != "" && ent.Level != PanicLevel {
		final.buf.AppendByte(' ')
		final.addAttribute(attributeFgRed)
		final.buf.AppendString("stacktrace")
		final.addAttribute(attributeReset)
		final.buf.AppendString("=\n")
		final.buf.AppendString(ent.Stack)
	}

	final.buf.AppendByte('\n')

	ret := final.buf
	putEncoder(final)
	return ret, nil
}

func init() {
	zap.RegisterEncoder("pretty", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewPrettyEncoder(cfg), nil
	})
}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

func (enc *prettyEncoder) addElementSeparator() {
	last := enc.buf.Len() - 1
	if last < 0 {
		return
	}
	switch enc.buf.Bytes()[last] {
	case '{', '[', ':', ',', '=', ' ':
		return
	default:
		enc.buf.AppendByte(',')
		enc.buf.AppendByte(' ')
	}
}

func (enc *prettyEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	enc.addKey(key)
	return enc.AppendArray(arr)
}

func (enc *prettyEncoder) AppendArray(arr zapcore.ArrayMarshaler) error {
	enc.addElementSeparator()
	enc.buf.AppendByte('[')
	err := arr.MarshalLogArray(enc)
	enc.buf.AppendByte(']')
	return err
}

func (enc *prettyEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	enc.addKey(key)
	return enc.AppendObject(obj)
}

func (enc *prettyEncoder) AppendObject(obj zapcore.ObjectMarshaler) error {
	enc.addElementSeparator()
	enc.buf.AppendByte('{')
	err := obj.MarshalLogObject(enc)
	enc.buf.AppendByte('}')
	return err
}

func (enc *prettyEncoder) AddBinary(key string, value []byte) {
	enc.addKey(key)
	enc.AppendBinary(value)
}

func (enc *prettyEncoder) AppendBinary(value []byte) {
	enc.addElementSeparator()
	fmt.Fprintf(enc.buf, "0x%x", value)
}
func (enc *prettyEncoder) AddByteString(key string, value []byte) {
	enc.AddString(key, string(value))
}
func (enc *prettyEncoder) AppendByteString(value []byte) {
	enc.AppendString(string(value))
}
func (enc *prettyEncoder) AddBool(key string, value bool) {
	enc.addKey(key)
	enc.AppendBool(value)
}

func (enc *prettyEncoder) AppendBool(value bool) {
	enc.addElementSeparator()
	if value {
		enc.buf.AppendString("true")
	} else {
		enc.buf.AppendString("false")
	}
}

func (enc *prettyEncoder) AddComplex128(key string, value complex128) { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendComplex128(value complex128)          { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddComplex64(key string, value complex64)   { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendComplex64(value complex64)            { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddDuration(key string, value time.Duration) {
	enc.addKey(key)
	enc.AppendDuration(value)
}

func (enc *prettyEncoder) AppendDuration(value time.Duration) {
	enc.buf.AppendString(value.String())
}

func (enc *prettyEncoder) AddFloat64(key string, value float64) { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendFloat64(value float64)          { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddFloat32(key string, value float32) { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendFloat32(value float32)          { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddInt(key string, value int)         { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendInt(value int)                  { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddInt64(key string, value int64)     { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendInt64(value int64)              { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddInt32(key string, value int32)     { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendInt32(value int32)              { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddInt16(key string, value int16)     { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendInt16(value int16)              { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddInt8(key string, value int8)       { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendInt8(value int8)                { enc.AppendReflected(value) }

func (enc *prettyEncoder) AddUint(key string, value uint)       { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUint(value uint)                { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddUint64(key string, value uint64)   { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUint64(value uint64)            { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddUint32(key string, value uint32)   { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUint32(value uint32)            { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddUint16(key string, value uint16)   { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUint16(value uint16)            { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddUint8(key string, value uint8)     { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUint8(value uint8)              { enc.AppendReflected(value) }
func (enc *prettyEncoder) AddUintptr(key string, value uintptr) { enc.AddReflected(key, value) }
func (enc *prettyEncoder) AppendUintptr(value uintptr)          { enc.AppendReflected(value) }

func (enc *prettyEncoder) AddString(key, value string) {
	enc.addKey(key)
	enc.AppendString(value)
}

func (enc *prettyEncoder) AppendString(val string) {
	enc.addElementSeparator()
	enc.buf.AppendByte('"')
	for i := 0; i < len(val); i++ {
		b := val[i]
		switch b {
		case '\n':
			enc.buf.AppendString(`\n`)
		case '\t':
			enc.buf.AppendString(`\t`)
		case '\r':
			enc.buf.AppendString(`\r`)
		default:
			enc.buf.AppendByte(val[i])
		}
	}
	enc.buf.AppendByte('"')
}

func (enc *prettyEncoder) AddTime(key string, value time.Time) {
	enc.addKey(key)
	enc.AppendTime(value)
}

func (enc *prettyEncoder) AppendTime(value time.Time) {
	enc.buf.AppendString(value.Format(time.RFC3339))
}

// AddReflected uses reflection to serialize arbitrary objects, so it can be
// slow and allocation-heavy.
func (enc *prettyEncoder) AddReflected(key string, value interface{}) error {
	enc.addKey(key)
	return enc.AppendReflected(value)
}

func (enc *prettyEncoder) AppendReflected(value interface{}) error {
	enc.addElementSeparator()
	fmt.Fprintf(enc.buf, "%#v", value)
	return nil
}

// OpenNamespace opens an isolated namespace where all subsequent fields will
// be added. Applications can use namespaces to prevent key collisions when
// injecting loggers into sub-components or third-party libraries.
func (enc *prettyEncoder) OpenNamespace(key string) {}

const escape = "\x1b"

type attribute int

// Base attributes
const (
	attributeReset attribute = iota
)

// Foreground text colors
//nolint:deadcode,varcheck
const (
	attributeFgBlack attribute = iota + 30
	attributeFgRed
	attributeFgGreen
	attributeFgYellow
	attributeFgBlue
	attributeFgMagenta
	attributeFgCyan
	attributeFgWhite
)

// Foreground Hi-Intensity text colors
//nolint:deadcode,varcheck
const (
	attributeFgHiBlack attribute = iota + 90
	attributeFgHiRed
	attributeFgHiGreen
	attributeFgHiYellow
	attributeFgHiBlue
	attributeFgHiMagenta
	attributeFgHiCyan
	attributeFgHiWhite
)

// Background text colors
//nolint:deadcode,varcheck
const (
	attributeBgBlack attribute = iota + 40
	attributeBgRed
	attributeBgGreen
	attributeBgYellow
	attributeBgBlue
	attributeBgMagenta
	attributeBgCyan
	attributeBgWhite
)

// Background Hi-Intensity text colors
//nolint:deadcode,varcheck
const (
	attributeBgHiBlack attribute = iota + 100
	attributeBgHiRed
	attributeBgHiGreen
	attributeBgHiYellow
	attributeBgHiBlue
	attributeBgHiMagenta
	attributeBgHiCyan
	attributeBgHiWhite
)
