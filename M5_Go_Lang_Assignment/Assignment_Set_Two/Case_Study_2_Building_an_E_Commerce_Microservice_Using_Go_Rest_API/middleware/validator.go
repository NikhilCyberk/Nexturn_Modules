package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

func ValidateJSON() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method == "POST" || c.Request.Method == "PUT" {
            if c.ContentType() != "application/json" {
                c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
                    "error": "Content-Type must be application/json",
                })
                return
            }
        }
        c.Next()
    }
}