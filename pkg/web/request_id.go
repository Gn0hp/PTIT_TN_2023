package web

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	contextRequestID string = "contextRequestID"
	headerRequestID  string = "X-Request-ID"
)

// GetRequestId retrieves request id from context
func GetRequestId(c *gin.Context) string {
	if rId, ok := c.Get(contextRequestID); ok {
		return rId.(string)
	}
	return ""
}
func attachRequestId(c *gin.Context) {
	if rId := c.Request.Header.Get(headerRequestID); len(rId) > 0 {
		c.Set(contextRequestID, rId)
		return
	}
	c.Set(contextRequestID, uuid.New().String())
}

// RequestIdMiddleware is a middleware that injects UUID to each web request
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		attachRequestId(c)
		c.Writer.Header().Set(headerRequestID, GetRequestId(c))
		c.Next()
	}
}
