package log

import "go.uber.org/zap"

type SlackGoAdapter struct {
	zl *zap.SugaredLogger
}

// NewSlackGoAdapter is a wrapper around a zap.SugaredLogger that implements the
// slack.Logger interface for the github.com/slack-go/slack package.
func NewSlackGoAdapter(zl *zap.SugaredLogger) *SlackGoAdapter {
	return &SlackGoAdapter{
		zl: zl,
	}
}

func (l SlackGoAdapter) Output(calldepth int, s string) error {
	l.zl.Info(s)
	return nil
}
