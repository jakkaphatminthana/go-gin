package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			custom.Error(c, lastErr)
			c.Abort()
		}
	}
}
