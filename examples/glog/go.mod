module glog_example

go 1.17

require (
	github.com/golang/glog v1.0.0
	github.com/planetscale/log v0.0.0-20211101173951-39b30043a122
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1
)

require github.com/slok/noglog v0.2.0 // indirect

replace github.com/golang/glog => github.com/planetscale/noglog v0.2.1-0.20210421230640-bea75fcd2e8e
