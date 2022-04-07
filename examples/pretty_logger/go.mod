module pretty_logger_example

go 1.18

replace github.com/planetscale/log => ../../

require github.com/planetscale/log v0.0.0-00010101000000-000000000000

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
)
