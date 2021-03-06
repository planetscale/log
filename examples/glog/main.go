package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/planetscale/log"
)

// Add to `go.mod`:
//
//   replace github.com/google/glog => github.com/slok/noglog master

func main() {
	logger := log.New()
	defer logger.Sync()

	// configure glog to log info to stdout, for demo purposes
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()
	defer glog.Flush()

	// hijack glog's logger and redirect it through zap logger
	slogger := logger.Sugar()
	glog.SetLogger(&glog.LoggerFunc{
		DebugfFunc: func(f string, a ...interface{}) { slogger.Debugf(f, a...) },
		InfofFunc:  func(f string, a ...interface{}) { slogger.Infof(f, a...) },
		WarnfFunc:  func(f string, a ...interface{}) { slogger.Warnf(f, a...) },
		ErrorfFunc: func(f string, a ...interface{}) { slogger.Errorf(f, a...) },
	})

	// zap logger:
	logger.Info("regular zap log")

	// hijacked glog:
	glog.Info("glog log message redirected to zap")
}
