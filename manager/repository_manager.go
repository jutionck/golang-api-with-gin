package manager

import (
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	productDb []model.Product
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(&r.productDb)
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
