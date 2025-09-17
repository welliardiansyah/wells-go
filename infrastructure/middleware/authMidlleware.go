package middleware

import (
	"net/http"
	"strings"
	"time"
	"wells-go/infrastructure/redis"
	"wells-go/response"
	"wells-go/util/security"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(maker security.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Authorization header missing", nil, startedAt)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid authorization header format", nil, startedAt)
			c.Abort()
			return
		}

		tokenStr := parts[1]

		val, err := redis.Rdb.Get(redis.Ctx, "jwt:"+tokenStr).Result()
		if err != nil || val != "active" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Token revoked or not found in cache", nil, startedAt)
			c.Abort()
			return
		}

		payload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid or expired token", err.Error(), startedAt)
			c.Abort()
			return
		}

		c.Set("user_id", payload.UserID)
		c.Set("roles", payload.Roles)
		c.Set("permissions", payload.Permissions)

		c.Next()
	}
}
