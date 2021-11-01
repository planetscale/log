# Examples

Runnable examples:

```console
go run ./logger.go
go run ./sugar_logger.go
...
```

Running the `glog.go` example requires additional steps:

1. Un-comment or add to `go.mod`:

```go
replace github.com/golang/glog => github.com/planetscale/noglog v0.2.1-0.20210421230640-bea75fcd2e8e
```

2. Run or build with `-build hijack_glog`  flag:

```console
go run -tags hijack_glog ./examples/glog.go
```

Run with `PS_DEV_MODE` env var set to show human friendly, non-JSON logging mode:

```console
PS_DEV_MODE=1 go run ./logger.go
```