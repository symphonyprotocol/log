# log
Symphonyprotocol logging module

## Status
`In development`, try to use branch `dev`, for now only console log supported.

## Usage
```go
import "github.com/symphonyprotocol/log"

var logger = log.GetLogger("Test Category")
logger.Debug("Hello %v", "World")
logger.Trace("Hello World")
```
