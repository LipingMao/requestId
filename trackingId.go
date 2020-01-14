package trackingId

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const DefaultHeader = "Tracking-Id"

func NewTrackingId() string {
	return uuid.New().String()
}

// Gin Middleware with default header
func TrackingId() gin.HandlerFunc {
	return TrackingIdWithCustomizedHeader(DefaultHeader)
}

// Gin Middleware with cusomized header
func TrackingIdWithCustomizedHeader(head string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(head)
		// Generate TrackingID if not exist
		if tId == "" {
			tId = NewTrackingId()
			c.Header(head, tId)
		}

		// Set in Context
		c.Set(head, tId)
		c.Next()
	}
}
