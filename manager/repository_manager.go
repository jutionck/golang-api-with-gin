package manager

import "github.com/jutionck/golang-api-with-gin/repository"

type RepositoryManager interface {
	// Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct{}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository()
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
