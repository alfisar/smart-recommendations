package service

import (
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"
)

// contract function for service product
type ProductSerivceContract interface {
	/*
		Get data recomendation product
		if data exists on redis then get data from redis
		else get data from bigtable and store to redis
	*/
	Get(key string) (result domain.Product, err errorhandler.ErrorData)
	// Insert data to bigtable from csv
	Insert() (err errorhandler.ErrorData)
}
