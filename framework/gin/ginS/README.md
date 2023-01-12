# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/realjunshi/beide/framework/gin"
	"github.com/realjunshi/beide/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
