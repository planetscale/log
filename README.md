# log 🪵

PlanetScale's logger module. Implemented as a set of convenience functions providing a common configuration of the [zap](https://github.com/uber-go/zap) logging library.  `zap` *"provides fast, structured, leveled logging."*

Logs emitted follow the standards from [PlanetScale Structure Logging (coda)](https://coda.io/d/Text-Book_dqagsnmBFI6/WIP-Structured-Logging_suS7P#_luvPS).

## Usage

```console
go get github.com/planetscale/log
```

`zap` provides two logging interfaces - The standard `zap.Logger` and the `zap.SugarLogger`. Zap describes each logger and when to choose between them as:

> In contexts where performance is nice, but not critical, use the `zap.SugaredLogger`. It's 4-10x faster than other structured logging packages and supports both structured and printf-style logging. Like log15 and go-kit, the SugaredLogger's structured logging APIs are loosely typed and accept a variadic number of key-value pairs. (For more advanced use cases, they also accept strongly typed fields - see the [SugaredLogger.With](https://pkg.go.dev/go.uber.org/zap@v1.19.1#SugaredLogger.With) documentation for details.

> In the rare contexts where every microsecond and every allocation matter, use the `zap.Logger`. It's even faster than the SugaredLogger and allocates far less, but it only supports strongly-typed, structured logging.

### Examples:

**zap.Logger**:

```go
import "github.com/planetscale/log"

func main() {
  fields := zap.Fields(zap.String("app", "logging-demo"))

  logger, _ := log.NewPlanetScaleLogger(fields)
  defer logger.Sync()

  logger.Info("info log with fields",
    // Structured context as typed key-value pairs
    zap.String("user_id", "12345678"),
    zap.String("branch_id", "xzyhnkhpi12"),
  )
}
```

**zap.SugarLogger**:

```go
import "github.com/planetscale/log"

func main() {
  fields := zap.Fields(zap.String("app", "logging-demo"))

  logger, _ := log.NewPlanetScaleSugarLogger(fields)
  defer logger.Sync()

  logger.Infof("sugar log printf example: %v", "foo")

  logger.Infow("sugar log with fields",
    // Structured context as loosely typed key-value pairs.
    "user_id", "12345678",
    "branch_id", "xzyhnkhpi12",
  )
}
```

See [./examples](./examples) for more.

### glog

Many PlanetScale applications use the [github.com/golang/glog](https://github.com/golang/glog) library which is commonly used in Vitess and Kuberenetes client libraries. Glog has some interesting properties, namely that it hooks into `flags` for configuration and outputs its own log messages. When combined with `zap` you will end up with an application that is mixing structured JSON logs with plain-text logs from `glog`.

Using [noglog](https://github.com/planetscale/noglog) the `glog` library's log calls can be replaced with our logger such that all logs emitted by the application are in a common, structured, JSON format.

1. Add the following to your `go.mod`:

```golang
require (
    github.com/google/glog master
)

replace github.com/google/glog => github.com/planetscale/noglog master
```

2. Replace `glog's` log calls with our SugaredLogger:

```golang
  logger, _ := log.NewPlanetScaleSugarLogger()
  defer logger.Sync()

  glog.SetLogger(&glog.LoggerFunc{
    DebugfFunc: func(f string, a ...interface{}) { logger.Debugf(f, a...) },
    InfofFunc:  func(f string, a ...interface{}) { logger.Infof(f, a...) },
    WarnfFunc:  func(f string, a ...interface{}) { logger.Warnf(f, a...) },
    ErrorfFunc: func(f string, a ...interface{}) { logger.Errorf(f, a...) },
  })
```

If using the `zap.Logger`:

```golang
  logger, _ := log.NewPlanetScaleLogger()
  defer logger.Sync()

  slogger := logger.Sugar()
  glog.SetLogger(&glog.LoggerFunc{
    DebugfFunc: func(f string, a ...interface{}) { slogger.Debugf(f, a...) },
    InfofFunc:  func(f string, a ...interface{}) { slogger.Infof(f, a...) },
    WarnfFunc:  func(f string, a ...interface{}) { slogger.Warnf(f, a...) },
    ErrorfFunc: func(f string, a ...interface{}) { slogger.Errorf(f, a...) },
  })

```

## Development mode

All logs are emitted as JSON by default. Sometimes this can be difficult to read. Set the `PS_DEV_MODE=1` environment variable to switch into a more human friendly log format.
