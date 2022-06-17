package main

import (
	"errors"
	"time"

	"github.com/planetscale/log"
)

type S struct {
	a string
}

func main() {
	logger, _ := log.NewPlanetScaleConfig("pretty", log.DebugLevel).Build()
	defer logger.Sync()

	logger.Debug("debugging")
	logger.Info("hello", log.String("foo", "bar"))
	time.Sleep(250 * time.Millisecond)
	logger.Info("test")

	logger = logger.Named("main")
	logger.Warn("ohno",
		log.Any("any-nil", nil),
		log.Any("any-struct", S{"hi"}),
		log.Any("any-string", "foo"),
		log.Any("any-int", 1),
		log.Binary("binary", []byte{0, 1, 2, 3, 12}),
		log.Int("int", 10),
		log.Int64("int64", 10),
		log.Bool("bool", true),
		log.Time("time", time.Now()),
		log.Duration("duration", 1*time.Second),
		log.String("string", "foo\nbar"),
		log.Strings("strings", []string{"a", "b"}),
		log.Bools("bools", []bool{true, false}),
		log.Stack("stack"),
		log.Uint16("uint16", 100),
		log.Float32("float32", 1.1),
	)

	time.Sleep(500 * time.Millisecond)

	logger.With(log.String("foo", "bar")).With(log.String("a", "b")).Info("hello", log.Int("int", 1))
	logger.With(log.Error(errors.New("bye"))).Error("wups 1")
	logger.Error("wups", log.Error(errors.New("bye")))
	logger.Fatal("test")
	logger.Panic("panic")
}
