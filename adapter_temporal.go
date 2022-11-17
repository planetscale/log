package log

import "go.uber.org/zap"

type TemporalAdapter struct {
	zl *zap.SugaredLogger
}

// NewTemporalAdapter is a wrapper around a zap.SugaredLogger that implements the
// temporal client.Logger interface for the github.com/temporalio/sdk-go package.
func NewTemporalAdapter(zl *zap.SugaredLogger) *TemporalAdapter {
	return &TemporalAdapter{
		zl: zl,
	}
}

func (l *TemporalAdapter) Debug(msg string, keyvals ...interface{}) {
	l.zl.Debugw(msg, keyvals...)
}

func (l *TemporalAdapter) Info(msg string, keyvals ...interface{}) {
	l.zl.Infow(msg, keyvals...)
}

func (l *TemporalAdapter) Warn(msg string, keyvals ...interface{}) {
	l.zl.Warnw(msg, keyvals...)
}

func (l *TemporalAdapter) Error(msg string, keyvals ...interface{}) {
	l.zl.Errorw(msg, keyvals...)
}
