package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/core/logger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			logger.Debug("OPTIONS request")
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// cors.New(cors.Config{
// 	AllowOrigins:     []string{"http://localhost:5173"}, // 允许的域名
// 	AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // 允许的 HTTP 方法
// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // 允许的请求头
// 	ExposeHeaders:    []string{"Content-Length"}, // 前端可以访问的响应头
// 	AllowCredentials: true, // 是否允许携带 Cookie
// 	MaxAge:           12 * 60 * 60, // 预检请求缓存时间（秒）
// }))
