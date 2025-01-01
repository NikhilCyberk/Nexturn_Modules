package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Missing or invalid Authorization header")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Decode the Base64-encoded credentials
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Failed to decode Authorization header:", err)
			c.JSON(401, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		// Split the username and password
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format:", string(payload))
			c.JSON(401, gin.H{"error": "Invalid Credentials"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		// Validate credentials against the database
		var storedPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&storedPassword)
		if err != nil {
			fmt.Println("User not found or error querying database:", err)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if storedPassword != password {
			fmt.Println("Password mismatch")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("Authentication successful")
		c.Next()
	}
}
