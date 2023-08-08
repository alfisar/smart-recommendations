package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	getRouterVersion(api)
	getRouterCredential(api)
	getRouterCProduct(api)
	return router
}
