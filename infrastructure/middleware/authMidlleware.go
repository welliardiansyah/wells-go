package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"wells-go/infrastructure/redis"
	"wells-go/response"
	"wells-go/util/security"
)

func AuthMiddleware(maker security.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Authorization header missing", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid authorization header format", nil)
			c.Abort()
			return
		}

		tokenStr := parts[1]

		val, err := redis.Rdb.Get(context.Background(), "jwt:"+tokenStr).Result()
		if err != nil || val != "active" {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Token revoked or not found in cache", nil)
			c.Abort()
			return
		}

		payload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		if payload.Roles == nil || len(payload.Roles) == 0 {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "roles not found in token", nil)
			c.Abort()
			return
		}

		c.Set(security.AuthorizationPayloadKey, payload)
		c.Set("user_id", payload.UserID)
		c.Set("roles", payload.Roles)
		c.Set("permissions", payload.Permissions)
		c.Next()
	}
}
