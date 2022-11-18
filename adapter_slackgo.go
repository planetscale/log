package log

import "go.uber.org/zap"

type SlackGoAdapter struct {
	zl *zap.Logger
}

// NewSlackGoAdapter is a wrapper around a zap.Logger that implements the
// slack.Logger interface for the github.com/slack-go/slack package.
func NewSlackGoAdapter(zl *zap.Logger) *SlackGoAdapter {
	return &SlackGoAdapter{
		// Skip one call frame to exclude zap_adapter itself.
		// Or it can be configured when logger is created (not always possible).
		zl: zl.WithOptions(zap.AddCallerSkip(1)),
	}
}

// Output implements the slack.Logger interface
func (l SlackGoAdapter) Output(calldepth int, s string) error {
	l.zl.Info(s)
	return nil
}
