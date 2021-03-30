# Bottom formatter for [logrus](https://github.com/sirupsen/logrus) based on [bottom-go](https://github.com/bottom-software-foundation/bottom-go)

## Usage

```go
package main

import (
        log "github.com/sirupsen/logrus"

        "github.com/hairmare/logrus-bottom/bottomlog"
)

func main() {

        // 🥺,,
        log.SetFormatter(bottomlog.NewBottomFormatter())

        log.WithFields(log.Fields{
                "animal": "walrus",
        }).Info("A walrus appears")
}
```

## Acknowledgements

The logrus-bottom team would like to thank the bottom software foundation for their inspiration.
