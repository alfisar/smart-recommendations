package controller

import (
	"smart-recommendation/internal/response"

	"github.com/gin-gonic/gin"
)

type ControllerVersion struct {
}

func NewControllerVersion() ControllerVersion {
	return ControllerVersion{}
}

func (obj ControllerVersion) GetVersion(ctx *gin.Context) {
	response.SuccessVersionOne(ctx.Writer)
	return
}
