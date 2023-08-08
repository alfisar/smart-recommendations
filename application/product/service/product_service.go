package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	repository "smart-recommendation/application/product/repository/bigtable"
	repositoryRedis "smart-recommendation/application/redis/repository"
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"
)

// struct to store the package repository that will be used
type ProductService struct {
	repo      repository.ProductRepositoryContract
	repoRedis repositoryRedis.CredentialRepositoryContRedis
}

// init new service product
func NewProductService(repo repository.ProductRepositoryContract, repoRedis repositoryRedis.CredentialRepositoryContRedis) *ProductService {
	return &ProductService{
		repo:      repo,
		repoRedis: repoRedis,
	}
}

// Get data recomendation product
// if data exists on redis then get data from redis
// else get data from bigtable and store to redis
func (obj *ProductService) Get(key string) (result domain.Product, err errorhandler.ErrorData) {
	redisKey := "rekomendation_produk_" + key
	data, _ := obj.repoRedis.Get(redisKey)
	if data != "" {
		errData := json.Unmarshal([]byte(data), &result)
		if errData != nil {
			errData := fmt.Errorf("Terjadi kesalahan - SM-10003")
			err = errorhandler.ErrParsingData("10003", errData)
			return
		}

		return
	}
	result, err = obj.repo.Get(key)
	if err.ResponseCode != "" {
		return
	}

	dataByte, errData := json.Marshal(&result)
	if errData != nil {
		errData := fmt.Errorf("Terjadi kesalahan - SM-10003")
		err = errorhandler.ErrParsingData("10003", errData)
		return
	}

	_ = obj.repoRedis.Insert(redisKey, string(dataByte))
	return
}

// Insert data to bigtable from csv
func (obj *ProductService) Insert() (err errorhandler.ErrorData) {
	file, errData := os.Open("MOCK_DATA.csv")
	if errData != nil {
		err = errorhandler.ErrValidationService("10002", fmt.Errorf("terjadi kesalahan - SB-10002"))
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, errData := reader.Read()
		if errData != nil {
			break
		}
		err = obj.repo.Insert("norek", "rekomendation", record[0], record[1])
		if err.ResponseCode != "" {
			return
		}

	}
	return
}
