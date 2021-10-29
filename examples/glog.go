package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/planetscale/log"
	"go.uber.org/zap"
)

func main() {
	// most services should set a global field `app=NAME` so we know who is talking
	fields := zap.Fields(zap.String("app", "logging-demo"))

	// setup zap
	logger, _ := log.NewPlanetScaleLogger(fields)
	defer logger.Sync()

	// setup glog for demo purposes
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()
	defer glog.Flush()

	// hijack glog's logger and redirect it through zap logger
	log.HijackGlog(logger)

	// zap logger:
	logger.Info("regular zap log")

	// hijacked glog:
	glog.Info("glog log message redirected to zap")

}
