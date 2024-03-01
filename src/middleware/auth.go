package middleware

import (
	"api-gin/src/models/user"
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims represents the JWT claims
type Claims struct {
	ID        string    `json:"id,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// You can add more claims as needed
	jwt.StandardClaims
}

// JWTMiddleware checks if the JWT token is valid
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Check if the Authorization header is in the format: Bearer <token>
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Parse the token from the Authorization header
		tokenString := authHeaderParts[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token is valid, extract user information from claims and set it in the context
		ctx := context.WithValue(c.Request.Context(), "user", &Claims{
			ID:        claims.ID,
			UserName:  claims.UserName,
			Gender:    claims.Gender,
			Active:    claims.Active,
			CreatedAt: claims.CreatedAt,
			UpdatedAt: claims.UpdatedAt,
		})
		c.Request = c.Request.WithContext(ctx)

		// Proceed with the request
		c.Next()
	}
}

// GenerateToken generates a JWT token
func GenerateToken(userData *user.Users) (string, error) {

	// Create the claims
	claims := &Claims{
		ID:        userData.ID.Hex(),
		UserName:  userData.UserName,
		Gender:    userData.Gender,
		Active:    userData.Active,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			// You can add more standard claims as needed
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
