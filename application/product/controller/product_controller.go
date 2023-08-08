package controller

import (
	"smart-recommendation/application/product/service"
	"smart-recommendation/internal/response"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductSerivceContract
}

func NewProductController(service service.ProductSerivceContract) ProductController {
	return ProductController{
		service: service,
	}
}

func (obj *ProductController) Get(ctx *gin.Context) {
	key := ctx.Param("norek")

	data, err := obj.service.Get(key)
	if err.ResponseCode != "" {
		response.Failed(ctx.Writer, err)
		return
	}
	response.Success(ctx.Writer, data)
	return
}

func (obj *ProductController) Insert(ctx *gin.Context) {
	err := obj.service.Insert()
	if err.ResponseCode != "" {
		response.Failed(ctx.Writer, err)
		return
	}
	response.Success(ctx.Writer, nil)
	return
}
