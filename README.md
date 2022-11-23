# wrzap

wrap zap logger

对zap日志库的包装，通过导入包的方式使logger全局唯一

- wrzap = wrap zap
- 无糖Logger

参考： https://github.com/bigwhite/experiments/tree/master/uber-zap-advanced-usage

## Dependency

go.uber.org/zap v1.23.0

## Usage

```go
package main

import log "github.com/micplus/wrzap"

func main() {
    defer log.Sync()
    log.Info("hello, zap",
        log.String("user", "myname"),
        log.Int("count", 1),
    )
}
```
