package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)

		// Process the request
		c.Next()

		// Log the response time
		fmt.Printf("Completed %s in %v\n", c.Request.URL.Path, time.Since(start))
	}
}
