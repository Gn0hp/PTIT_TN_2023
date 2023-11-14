package utils

import (
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Error     string      `json:"error,omitempty"`
	Data      interface{} `json:"data"`
}

func ResponseMiddleware(logger logur.LoggerFacade) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if skipGenResponse(c.Request.URL.Path) {
			return
		}
		c.Header("Content-Type", "application/json")
		if len(c.Errors) > 0 {
			c.JSON(http.StatusBadRequest, Response{
				Status:    http.StatusBadRequest,
				Timestamp: time.Now().Unix(),
				Error:     c.Errors[0].Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			Status:    http.StatusOK,
			Timestamp: time.Now().Unix(),
			Data:      c.Keys["data"],
		})
	}
}

func SetResponse(c *gin.Context, data interface{}) {
	c.Set("data", data)
}

func skipGenResponse(path string) bool {
	skipPath := []string{
		"/swagger",
	}
	for _, s := range skipPath {
		if strings.Contains(path, s) {
			return true
		}
	}
	return false
}
