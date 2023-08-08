package repository

import (
	"smart-recommendation/internal/errorhandler"

	"github.com/go-redis/redis"
)

type CredentialRepositoryRedis struct {
	conn *redis.Client
}

func NewCredentialRedis(conn *redis.Client) *CredentialRepositoryRedis {
	return &CredentialRepositoryRedis{
		conn: conn,
	}
}

func (obj *CredentialRepositoryRedis) Get(key string) (data string, err errorhandler.ErrorData) {
	data, errData := obj.conn.Get(key).Result()
	if errData != nil {
		err = errorhandler.ErrorRepo(errData)
	}
	return
}

func (obj *CredentialRepositoryRedis) Insert(key string, data string) (err errorhandler.ErrorData) {
	result := obj.conn.Set(key, data, 0)
	if result.Err() != nil {
		err = errorhandler.ErrorRepo(result.Err())
	}

	return
}
