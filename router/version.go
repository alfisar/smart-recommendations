package router

import "github.com/gin-gonic/gin"

func getRouterVersion(router *gin.RouterGroup) {
	controller := setVersionController()
	router.GET("", controller.GetVersion)
}
