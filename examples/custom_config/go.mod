module custom_config_example

go 1.18

require github.com/planetscale/log v0.0.0-20211222231218-1db482fc5936

replace github.com/planetscale/log => ../../

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)
