package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JWTConfig holds the configuration for JWT middleware
type JWTConfig struct {
    SecretKey []byte
}

// JWTAuth creates a JWT authentication middleware with the provided configuration
func JWTAuth(config JWTConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract token from Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "missing or invalid authorization header",
            })
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate the signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, errors.New("invalid signing method")
            }
            return config.SecretKey, nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid token: " + err.Error(),
            })
            return
        }

        if !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "token is not valid",
            })
            return
        }

        // Extract and validate claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid token claims",
            })
            return
        }

        // Validate required claims
        userID, ok := claims["user_id"]
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "missing user_id claim",
            })
            return
        }

        // Add claims to context for use in handlers
        c.Set("userID", userID)
        if role, exists := claims["role"]; exists {
            c.Set("role", role)
        }

        c.Next()
    }
}