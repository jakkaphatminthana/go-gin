package custom

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jakkaphatminthana/go-gin/utils"
)

type (
	ErrorResponse struct {
		Error ErrorDetail `json:"error"`
	}

	ErrorDetail struct {
		Status  int               `json:"status"`
		Message string            `json:"message"`
		Details map[string]string `json:"details,omitempty"`
	}
)

func Error(pctx *gin.Context, err error, statusCode ...int) {
	var (
		code    = http.StatusInternalServerError
		message = "Internal server error"
		details = make(map[string]string)
	)

	// case general AppError
	var appErr AppError
	if errors.As(err, &appErr) {
		code = appErr.StatusCode()
		message = appErr.Error()
	} else {
		code = statusCode[0]
		message = err.Error()
	}

	// case validation error (query, body)
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		code = http.StatusBadRequest
		message = "Validation failed"
		details = map[string]string{}
		for _, e := range validationErrs {
			jsonKey := utils.GetJsonFieldNameByErrorField(e)
			details[jsonKey] = validationErrorMessage(e)
		}
	}

	pctx.JSON(code, &ErrorResponse{
		Error: ErrorDetail{
			Status:  code,
			Message: message,
			Details: details,
		},
	},
	)
}

func validationErrorMessage(err validator.FieldError) string {
	jsonKey := utils.GetJsonFieldNameByErrorField(err)

	switch err.Tag() {
	case "required":
		return jsonKey + " is required"
	case "max":
		return jsonKey + " must be at most " + err.Param() + " characters"
	case "min":
		return jsonKey + " must be at least " + err.Param() + " characters"
	case "oneof":
		return jsonKey + " must be one of " + err.Param()
	default:
		return jsonKey + " is invalid"
	}
}
