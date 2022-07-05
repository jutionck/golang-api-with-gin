package usecase

import (
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
)

type ListProductUseCase interface {
	List() ([]model.Product, error)
}

type listProductUseCase struct {
	repo repository.ProductRepository
}

func (l *listProductUseCase) List() ([]model.Product, error) {
	return l.repo.Retrieve()
}

func NewListProductUseCase(repo repository.ProductRepository) ListProductUseCase {
	return &listProductUseCase{repo: repo}
}
