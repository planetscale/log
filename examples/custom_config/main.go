package main

import (
	"github.com/planetscale/log"
)

func main() {
	// disable the `caller` field in logs:
	cfg := log.NewPlanetScaleConfigDefault()
	logger, _ := cfg.Build(log.WithCaller(false))
	defer logger.Sync()

	logger.Info("basic log example")

	logger.Info("info log with fields",
		// Structured context as typed key-value pairs
		log.String("user_id", "12345678"),
		log.String("branch_id", "xzyhnkhpi12"),
	)

	logger.Warn("warning log with fields",
		log.String("user_id", "12345678"),
		log.String("branch_id", "xzyhnkhpi12"),
	)

	logger.Error("error log with fields",
		// Error logs will include a `stacktrace` field automatically
		log.String("user_id", "12345678"),
		log.String("branch_id", "xzyhnkhpi12"),
	)

	functionWithExtraContext(logger)
}

func functionWithExtraContext(logger *log.Logger) {
	// setup a temporarily logger with additional fields. All logs emitted from
	// this func will include the fields.
	l := logger.With(
		log.String("func", "extraContext"),
		log.String("transaction", "12345678"),
	)

	l.Info("logger with extra context")
	return
}
