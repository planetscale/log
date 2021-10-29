module github.com/planetscale/log

go 1.17

require (
	github.com/golang/glog v1.0.0
	github.com/slok/noglog v0.2.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1
)

// Enables hijacking glog. You also need to call `HijackGlog(logger)` where logger is a zap.Logger in your main()
replace github.com/golang/glog => github.com/planetscale/noglog v0.2.1-0.20210421230640-bea75fcd2e8e
