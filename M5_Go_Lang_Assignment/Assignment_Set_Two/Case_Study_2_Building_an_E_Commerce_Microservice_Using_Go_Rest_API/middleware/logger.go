package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start time
        startTime := time.Now()

        // Process request
        c.Next()

        // End time
        endTime := time.Now()

        // Execution time
        latencyTime := endTime.Sub(startTime)

        // Request method
        reqMethod := c.Request.Method

        // Request route
        reqUri := c.Request.RequestURI

        // Status code
        statusCode := c.Writer.Status()

        // Request IP
        clientIP := c.ClientIP()

        // Log format
        log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s",
            endTime.Format("2006/01/02 - 15:04:05"),
            statusCode,
            latencyTime,
            clientIP,
            reqMethod,
            reqUri,
        )
    }
}