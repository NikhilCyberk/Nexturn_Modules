package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login LoginRequest
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// In a real application, validate credentials against database
		// This is just for demonstration
		if login.Username != "admin" || login.Password != "password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = 1
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

		// Sign token
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}