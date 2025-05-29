package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ย้าย secretKey ไปใช้จาก environment variable
var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// VerifyJWTToken validates JWT token with the secret key
func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

// ExtractClaimsFromJWT extracts claims from JWT token
func ExtractClaimsFromJWT(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

// JWTMiddleware validates JWT token in request header
func JWTMiddleware() gin.HandlerFunc {
	// Add Line Numbers to Log Output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check if token exists and has Bearer prefix
		if tokenString == "" || len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Authorization header is required and must be Bearer token",
			})
			c.Abort()
			return
		}

		// Remove Bearer prefix
		tokenString = tokenString[7:]

		token, err := VerifyJWTToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		claims, ok := ExtractClaimsFromJWT(token)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Store claims in context for handlers to use
		c.Set("claims", claims)
		c.Next()
	}
}
