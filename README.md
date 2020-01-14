# trackingId
Gin trackingId Middleware, if the Request has Tracking-Id already, just store it in context, if it does not have Tracking-Id header, generate a new one with Google uuid.

# How to use

## With default header "Tracking-Id"
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
    r.Use(trackingId.TrackingId())

    // Your routes
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // Start server
    r.Run()
}
```

## With Customized Header

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
    r.Use(trackingId.TrackingIdWithCustomizedHeader("X-Track-ID"))

    // Your routes
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // Start server
    r.Run()
}
```
