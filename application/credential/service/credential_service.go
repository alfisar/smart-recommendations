package service

import (
	"fmt"
	"os"
	repository "smart-recommendation/application/credential/repository/mysql"
	repositoryRedis "smart-recommendation/application/redis/repository"
	"smart-recommendation/internal/errorhandler"
	hashGenerator "smart-recommendation/internal/hashgenerator"
	jwthandler "smart-recommendation/internal/jwtHandler"
	"time"
)

// struct to store the package repository that will be used
type CredentialService struct {
	repo      repository.CredentialRepositoryCont
	reporedis repositoryRedis.CredentialRepositoryContRedis
}

// init new service credential
func NewCredentialService(repo repository.CredentialRepositoryCont, reporedis repositoryRedis.CredentialRepositoryContRedis) *CredentialService {
	return &CredentialService{
		repo:      repo,
		reporedis: reporedis,
	}
}

func (obj *CredentialService) Login(application string, pass string) (token string, err errorhandler.ErrorData) {
	key := "TOKEN_" + application
	token, err = obj.reporedis.Get(key)
	if token != "" {
		err = errorhandler.ErrorData{}
		return
	}

	dataCred, err := obj.repo.GetByAplication(application)
	if err.ResponseCode != "" {
		return
	}
	// pass, err = hashGenerator.Generate(pass)
	valid, err := hashGenerator.Verify(dataCred.Password, pass)

	if !valid {
		err = errorhandler.ErrValidationService("1001", fmt.Errorf("terjadi kesalahan - SM-10001"))
		return
	}

	objJwt := jwthandler.JwtHandler{
		Secret: os.Getenv("JWT_SECRET"),
	}

	token, err = objJwt.GetToken(15 * time.Minute)
	obj.reporedis.Insert(key, token)
	return
}
