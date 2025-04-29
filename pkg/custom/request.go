package custom

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type (
	RequestBinder interface {
		Bind(obj any) error
	}

	requestBinder struct {
		ctx *gin.Context
	}
)

func NewCustomRequest(ctx *gin.Context) RequestBinder {
	return &requestBinder{ctx: ctx}
}

// implement
func (c *requestBinder) Bind(obj any) error {
	if err := c.ctx.ShouldBind(obj); err != nil {
		return err
	}

	if err := validate.Struct(obj); err != nil {
		return err
	}

	return nil
}
