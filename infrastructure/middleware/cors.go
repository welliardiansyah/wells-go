package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"wells-go/infrastructure/config"
)

func CORSMiddleware(config *config.Config) gin.HandlerFunc {
	allowedOrigins := strings.Split(config.AllowedOrigins, ",")
	allowedMethods := config.AllowedMethods
	env := strings.ToLower(config.Environment)

	fmt.Println("🚀 CORS Middleware initialized:")
	fmt.Println("   Environment:", env)
	fmt.Println("   Allowed origins:", allowedOrigins)
	fmt.Println("   Allowed methods:", allowedMethods)

	return func(c *gin.Context) {
		origin := strings.TrimSpace(c.Request.Header.Get("Origin"))
		if origin != "" {
			fmt.Printf("🌐 CORS Request: %s %s from %s\n", c.Request.Method, c.Request.URL.Path, origin)

			allowOrigin := ""
			if env == "development" || config.AllowedOrigins == "" {
				allowOrigin = "*"
				fmt.Println("   ✅ Using wildcard (development mode)")
			} else {
				for _, o := range allowedOrigins {
					trimmed := strings.TrimSpace(o)
					if origin == trimmed {
						allowOrigin = origin
						fmt.Println("   ✅ Origin matched:", trimmed)
						break
					}
				}
				if allowOrigin == "" {
					fmt.Printf("   ❌ Origin '%s' not in allowed list %v\n", origin, allowedOrigins)
				}
			}

			if allowOrigin != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				c.Writer.Header().Set("Access-Control-Allow-Headers",
					"Authorization, Content-Type, Content-Length, Accept, Accept-Encoding, "+
						"X-CSRF-Token, X-Requested-With, Cache-Control, Origin")
				if allowedMethods != "" {
					c.Writer.Header().Set("Access-Control-Allow-Methods", allowedMethods)
				} else {
					c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				}
				fmt.Println("   ✅ CORS headers set")
			}

			if c.Request.Method == "OPTIONS" {
				fmt.Println("   🔄 Handling OPTIONS preflight - returning 200")
				c.AbortWithStatus(200)
				return
			}
		}

		c.Next()
	}
}
