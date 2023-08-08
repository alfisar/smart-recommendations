package service

import (
	"smart-recommendation/internal/errorhandler"
)

// contract function service credential
type CredentialServiceCont interface {
	// Get data token from application name and password, if exist on redis get data form redis, if not exist get data from db and store to redis
	Login(application string, pass string) (token string, err errorhandler.ErrorData)
}
