package manager

import (
	"github.com/jutionck/golang-api-with-gin/repository"
)

type RepositoryManager interface {
	// ProductRepo Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
