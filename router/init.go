package router

import (
	"smart-recommendation/application/version/controller"
	_controller "smart-recommendation/application/version/controller"
	"smart-recommendation/database"

	_credController "smart-recommendation/application/credential/controller"
	_credRepoSQL "smart-recommendation/application/credential/repository/mysql"
	_credServ "smart-recommendation/application/credential/service"

	_redisRepo "smart-recommendation/application/redis/repository"

	_productController "smart-recommendation/application/product/controller"
	_productRepoBigtbl "smart-recommendation/application/product/repository/bigtable"
	_productServ "smart-recommendation/application/product/service"

	"github.com/joho/godotenv"
)

var (
	_         = godotenv.Load(".env")
	db        = database.NewDatabaseMySql()
	redis     = database.NewDatabaseRedis()
	bigtable  = database.NewBigtable()
	repoRedis = _redisRepo.NewCredentialRedis(redis)
)

func setVersionController() *controller.ControllerVersion {
	controller := _controller.NewControllerVersion()
	return &controller
}

func setCredentialController() *_credController.CredentialController {
	repo := _credRepoSQL.NewCredentialRepository(db)
	service := _credServ.NewCredentialService(repo, repoRedis)
	controller := _credController.NewCredentialController(service)
	return &controller
}

func setProductController() *_productController.ProductController {
	repo := _productRepoBigtbl.NewProductRepository(bigtable)
	service := _productServ.NewProductService(repo, repoRedis)
	controller := _productController.NewProductController(service)
	return &controller
}
