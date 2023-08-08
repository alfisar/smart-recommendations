package repository

import (
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"
)

// contract function for repository product
type ProductRepositoryContract interface {
	// Get data from bigtable using key
	Get(key string) (result domain.Product, err errorhandler.ErrorData)

	// Insert data to bigtable
	Insert(column string, columnTwo string, data string, datatwo string) (err errorhandler.ErrorData)
}
