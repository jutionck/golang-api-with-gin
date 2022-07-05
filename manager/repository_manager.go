package manager

import (
	"github.com/jutionck/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	productRepo repository.ProductRepository
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return r.productRepo
}

func NewRepositoryManager() RepositoryManager {
	productRepo := repository.NewProductRepository()
	return &repositoryManager{
		productRepo: productRepo,
	}
}
