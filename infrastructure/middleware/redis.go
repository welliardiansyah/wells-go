package middleware

import (
	"net/http"
	"wells-go/infrastructure/redis"
	"wells-go/response"

	"github.com/gin-gonic/gin"
)

func RedisMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := redis.Rdb.Ping(redis.Ctx).Result()
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Redis unavailable", err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
