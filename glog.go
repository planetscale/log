//go:build glog
// +build glog

package log

import (
	"github.com/golang/glog"
	"go.uber.org/zap"
)

func HijackGlog(logger *zap.Logger) {
	slogger := logger.Sugar()
	glog.SetLogger(&glog.LoggerFunc{
		DebugfFunc: func(f string, a ...interface{}) { slogger.Debugf(f, a...) },
		InfofFunc:  func(f string, a ...interface{}) { slogger.Infof(f, a...) },
		WarnfFunc:  func(f string, a ...interface{}) { slogger.Warnf(f, a...) },
		ErrorfFunc: func(f string, a ...interface{}) { slogger.Errorf(f, a...) },
	})
}
