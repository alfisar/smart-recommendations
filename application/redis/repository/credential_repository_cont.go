package repository

import (
	"smart-recommendation/internal/errorhandler"
)

type CredentialRepositoryContRedis interface {
	Get(key string) (data string, err errorhandler.ErrorData)
	Insert(key string, data string) (err errorhandler.ErrorData)
}
