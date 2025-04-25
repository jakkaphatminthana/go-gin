package custom

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	GinRequest interface {
		Bind(obj any) error
	}

	customGinRequest struct {
		ctx       *gin.Context
		validator *validator.Validate
	}
)

func NewCustomRequest(ginRequest *gin.Context) GinRequest {
	return &customGinRequest{
		ctx:       ginRequest,
		validator: validator.New(),
	}
}

// implement
func (c *customGinRequest) Bind(obj any) error {
	if err := c.ctx.ShouldBind(obj); err != nil {
		return err
	}

	if err := c.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
