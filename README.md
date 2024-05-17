# logger

Simple logger with standard output only.

install

```bash
go get github.com/riiim400th/logger
```

sample

```go
import	"github.com/riiim400th/logger"

logger.Log(logger.Debug, AnyType...)
logger.Log(logger.Info, AnyType...)
logger.Log(logger.Error, AnyType...)
//When specifying 'panic', the execution of Go stops.
logger.Log(logger.Panic, AnyType...)
```
