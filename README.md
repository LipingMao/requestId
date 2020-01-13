# trackingId
Gin trackingId Middleware

# How to use

```go
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/LipingMao/trackingId"
    ...
)

func main() {
    // Initialize router
    r := gin.Default()

    // Load middleware
    r.Use(trackingId.TrackingId)

    // Your routes
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // Start server
    r.Run()
}
```