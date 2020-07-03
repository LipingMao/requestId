package trackingId

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const defaultHeader = "Tracking-Id"

func GetTrackingId(c *gin.Context) string {
	return c.GetString(defaultHeader)
}

func GetHeaderName() string {
	return defaultHeader
}

// Gin Middleware with default header
func TrackingId() gin.HandlerFunc {
	return trackingIdWithCustomizedHeader(defaultHeader)
}

// Gin Middleware with cusomized header
func trackingIdWithCustomizedHeader(head string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(head)
		// Generate TrackingID if not exist
		if tId == "" {
			tId = uuid.New().String()
			c.Header(head, tId)
		}

		// Set in Context
		c.Set(head, tId)
		c.Next()
	}
}
