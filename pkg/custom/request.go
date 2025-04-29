package custom

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type (
	RequestBinder interface {
		BindBody(obj any) error
		BindQuery(obj any) error
		BindUri(obj any) error
	}

	requestBinder struct {
		ctx *gin.Context
	}
)

func NewCustomRequest(ctx *gin.Context) RequestBinder {
	return &requestBinder{ctx: ctx}
}

// implement
func (c *requestBinder) BindBody(obj any) error {
	if err := c.ctx.ShouldBind(obj); err != nil {
		return err
	}

	if err := validate.Struct(obj); err != nil {
		return err
	}

	return nil
}

// implement
func (c *requestBinder) BindQuery(obj any) error {
	if err := c.ctx.ShouldBindQuery(obj); err != nil {
		return err
	}
	return validate.Struct(obj)
}

// implement
func (c *requestBinder) BindUri(obj any) error {
	if err := c.ctx.ShouldBindUri(obj); err != nil {
		return err
	}
	return validate.Struct(obj)
}
