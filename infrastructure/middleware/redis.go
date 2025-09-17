package middleware

import (
	"net/http"
	"time"
	"wells-go/infrastructure/redis"
	"wells-go/response"

	"github.com/gin-gonic/gin"
)

func RedisMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()

		_, err := redis.Rdb.Ping(redis.Ctx).Result()
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Redis unavailable", err.Error(), startedAt)
			c.Abort()
			return
		}

		c.Next()
	}
}
