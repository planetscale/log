package main

import (
	"github.com/planetscale/log"
	"go.uber.org/zap"
)

// TODO copy rationale for using sugar vs regular logger here

func main() {
	// most services should set a global field `app=NAME` so we know who is talking
	fields := zap.Fields(zap.String("app", "sugar-logger"))

	logger, _ := log.NewPlanetScaleSugarLogger(fields)
	defer logger.Sync()

	logger.Info("basic sugar log example")
	logger.Infof("sugar log *f example: %v", "foo")
	logger.Infow("sugar log with fields",
		// Structured context as loosely typed key-value pairs.
		"user_id", "12345678",
		"branch_id", "xzyhnkhpi12",
	)
}
