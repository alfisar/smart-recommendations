package controller

import "github.com/gin-gonic/gin"

type CredentialControllerCont interface {
	login(ctx *gin.Context)
}
