package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// Data response
	ResponseSuccess struct {
		Data any `json:"data"`
	}

	// Meta response
	MetaSuccessResponse struct {
		Success SuccessDetail `json:"success"`
	}

	SuccessDetail struct {
		Status  int               `json:"status"`
		Name    string            `json:"name"`
		Message string            `json:"message"`
		Details map[string]string `json:"details,omitempty"`
	}
)

func Success(pctx *gin.Context, statusCode int, data any) {
	pctx.JSON(statusCode, &ResponseSuccess{Data: data})
}

func MetaSuccess(pctx *gin.Context, statusCode int, message string, details map[string]string) {
	if details == nil {
		details = make(map[string]string)
	}

	pctx.JSON(statusCode, &MetaSuccessResponse{
		Success: SuccessDetail{
			Status:  statusCode,
			Name:    http.StatusText(statusCode),
			Message: message,
			Details: details,
		},
	})
}
