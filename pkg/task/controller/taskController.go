package controller

import "github.com/gin-gonic/gin"

type TaskController interface {
	Listing(pctx *gin.Context)
	FindById(pctx *gin.Context)
}
