module github.com/planetscale/log

go 1.17

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1
)

// Enables hijacking glog
// replace github.com/golang/glog => github.com/slok/noglog v0.2.0
