package main

import (
	"github.com/planetscale/log"
)

func main() {
	logger := log.NewPlanetScaleSugarLogger()
	defer logger.Sync()

	logger.Info("basic info log example")
	logger.Infof("info log printf example: %v", "foo")
	logger.Infow("info log with fields",
		// Structured context as loosely typed key-value pairs.
		"user_id", "12345678",
		"branch_id", "xzyhnkhpi12",
	)
}
