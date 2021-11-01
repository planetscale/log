package main

import (
	"github.com/planetscale/log"
	"go.uber.org/zap"
)

// TODO copy differences between sugar logger and zap logger here, or maybe we just do that in the readme?

func main() {
	// most services should set a global field `app=NAME` so we know who is talking
	fields := zap.Fields(zap.String("app", "logging-demo"))

	// setup zap logger
	logger, _ := log.NewPlanetScaleLogger(fields)
	defer logger.Sync()

	logger.Info("basic log example")

	logger.Info("info log with fields",
		// Structured context as typed key-value pairs
		zap.String("user_id", "12345678"),
		zap.String("branch_id", "xzyhnkhpi12"),
	)

	logger.Warn("warning log with fields",
		zap.String("user_id", "12345678"),
		zap.String("branch_id", "xzyhnkhpi12"),
	)

	logger.Error("error log with fields",
		// Error logs will include a `stacktrace` field automatically
		zap.String("user_id", "12345678"),
		zap.String("branch_id", "xzyhnkhpi12"),
	)

	functionWithExtraContext(logger)
}

func functionWithExtraContext(logger *zap.Logger) {
	// setup a temporarily logger with additional fields. All logs emitted from
	// this func will include the fields.
	l := logger.With(
		zap.String("func", "extraContext"),
		zap.String("transaction", "12345678"),
	)

	l.Info("logger with extra context")
	return
}
