package custom

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErorrMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Error(pctx *gin.Context, err error, statusCode ...int) {
	var appErr AppError

	// AppError Handler
	if errors.As(err, &appErr) {
		pctx.JSON(
			appErr.StatusCode(),
			&ErorrMessage{
				StatusCode: appErr.StatusCode(),
				Message:    appErr.Error(),
			},
		)
		return
	}

	if len(statusCode) > 0 {
		pctx.JSON(
			statusCode[0],
			&ErorrMessage{
				StatusCode: statusCode[0],
				Message:    err.Error(),
			},
		)
		return
	}

	pctx.JSON(
		http.StatusInternalServerError,
		&ErorrMessage{StatusCode: http.StatusInternalServerError, Message: err.Error()},
	)
}
