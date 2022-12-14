// wrap zap logger
package wrzap

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3
	PanicLevel  Level = zap.PanicLevel  // 4
	FatalLevel  Level = zap.FatalLevel  // 5
	DebugLevel  Level = zap.DebugLevel  // -1
)

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

var (
	Skip        = zap.Skip
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
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float64s    = zap.Float64s
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Float32s    = zap.Float32s
	Int         = zap.Int
	Intp        = zap.Intp
	Ints        = zap.Ints
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int64s      = zap.Int64s
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int32s      = zap.Int32s
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int16s      = zap.Int16s
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	Int8s       = zap.Int8s
	String      = zap.String
	Stringp     = zap.Stringp
	Strings     = zap.Strings
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uints       = zap.Uints
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint64s     = zap.Uint64s
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint32s     = zap.Uint32s
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint16s     = zap.Uint16s
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uint8s      = zap.Uint8s
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Uintptrs    = zap.Uintptrs
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Times       = zap.Times
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Durations   = zap.Durations
	Any         = zap.Any

	Info   = logger.Info
	Warn   = logger.Warn
	Error  = logger.Error
	DPanic = logger.DPanic
	Panic  = logger.Panic
	Fatal  = logger.Fatal
	Debug  = logger.Debug
)

// ????????????logger???????????????logger???
// ?????????????????????????????????(???logger????????????)?????????????????????????????????????????????
func ResetDefault(l *Logger) {
	logger = l
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	DPanic = logger.DPanic
	Panic = logger.Panic
	Fatal = logger.Fatal
	Debug = logger.Debug
}

type Logger struct {
	// zap?????????????????????
	l     *zap.Logger
	level Level
}

// ????????????
var logger = New(os.Stderr, InfoLevel)

func Default() *Logger {
	return logger
}

func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("nil log writer")
	}
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		l:     zap.New(core),
		level: level,
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}
