package router

import "github.com/gin-gonic/gin"

func getRouterCredential(router *gin.RouterGroup) {
	controller := setCredentialController()
	router.POST("/login", controller.Login)
}
