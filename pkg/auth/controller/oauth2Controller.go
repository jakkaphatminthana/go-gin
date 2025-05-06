package controller

import (
	"github.com/gin-gonic/gin"
)

type OAuth2Controller interface {
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
}
