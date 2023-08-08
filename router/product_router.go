package router

import "github.com/gin-gonic/gin"

func getRouterCProduct(router *gin.RouterGroup) {
	controller := setProductController()
	router.GET("/product-recomendation/:norek", controller.Get)
	router.POST("/product-recomendation", controller.Insert)
}
