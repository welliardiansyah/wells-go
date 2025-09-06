package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wells-go/response"
	"wells-go/util/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(strings.ToLower(h), "bearer ") {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "missing bearer token", nil)
			c.Abort()
			return
		}

		token := strings.TrimSpace(h[7:])
		claims, err := jwt.Parse(token)
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "invalid or expired token", err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", claims.Sub)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}
