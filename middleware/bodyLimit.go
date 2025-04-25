package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/utils"
)

func BodyLimitMiddleware(bodyLimit string) gin.HandlerFunc {
	limitBytes := utils.ParseSize(bodyLimit)
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limitBytes)
		c.Next()
	}
}
