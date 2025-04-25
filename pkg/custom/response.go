package custom

import "github.com/gin-gonic/gin"

type ResponseSuccess struct {
	Data any `json:"data"`
}

func Success(pctx *gin.Context, statusCode int, data any) {
	pctx.JSON(statusCode, &ResponseSuccess{Data: data})
}
