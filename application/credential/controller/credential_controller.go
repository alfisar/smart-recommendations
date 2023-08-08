package controller

import (
	"smart-recommendation/application/credential/service"
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"
	"smart-recommendation/internal/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CredentialController struct {
	service service.CredentialServiceCont
}

func NewCredentialController(service service.CredentialServiceCont) CredentialController {
	return CredentialController{
		service: service,
	}
}

func (obj *CredentialController) Login(ctx *gin.Context) {
	var data domain.Credential
	errData := ctx.BindJSON(&data)
	if errData != nil {
		response.Failed(ctx.Writer, errorhandler.ErrController(errData))
		return
	}

	token, err := obj.service.Login(data.Application, data.Password)
	if err.ResponseCode != "" {
		response.Failed(ctx.Writer, err)
		return
	}

	result := domain.TokenLogin{
		Token:   token,
		Expired: strconv.Itoa(int(15 * time.Minute)),
	}

	response.Success(ctx.Writer, result)
	return
}
