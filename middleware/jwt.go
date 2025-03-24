package middleware

import (
	"crypto/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/model/system"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func init() {
	if len(secretKey) == 0 {
		// secretKey := make([]byte, 32) // 256 位密钥
		_, err := rand.Read(secretKey)
		if err != nil {
			logger.Fatal("Generate secret key error: %s", err)
		}
		// secretKey = key
		logger.Debug("Generate secret key: %s", secretKey)
	}
}

type Claims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJwt(user *system.SysUser) (string, error) {
	claims := Claims{
		UserName: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func parseJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		logger.Info("Parse token error: %s", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(401, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}
		// token 是 "Bearer <token>"
		bearerToken := strings.Split(token, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(400, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}
		tokenString := bearerToken[1]

		claims, err := parseJwt(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
