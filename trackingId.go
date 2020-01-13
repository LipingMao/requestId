package trackingId

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TrackingID = "Tracking-Id"

func NewTrackingId() string {
	return uuid.New().String()
}

// Gin Middleware
func TrackingId() gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(TrackingID)
		// Generate TrackingID is not exist
		if tId == "" {
			tId = NewTrackingId()
			c.Header(TrackingID, tId)
		}

		// Set in Context
		c.Set(TrackingID, tId)
		c.Next()
	}
}
