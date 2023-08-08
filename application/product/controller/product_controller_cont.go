package controller

import "github.com/gin-gonic/gin"

type ProductControlerContract interface {
	Get(ctx *gin.Context)
}
