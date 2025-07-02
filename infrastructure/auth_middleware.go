package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSvc = NewJwtService(os.Getenv("JWT_SECRET"))

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		token, err := jwtSvc.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		fmt.Printf("token.Valid: %v\n", token.Valid)
		fmt.Printf("token.Claims: %#v\n", token.Claims)

		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", (*claims)["user_id"])
			c.Set("username", (*claims)["username"])
			c.Set("first_name", (*claims)["first_name"])
			c.Set("last_name", (*claims)["last_name"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
