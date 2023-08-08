package controller

import "github.com/gin-gonic/gin"

type ControllerVersionContract interface {
	GetVersion(ctx *gin.Context)
}
