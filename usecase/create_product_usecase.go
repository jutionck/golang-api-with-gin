package usecase

import (
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product) error
}

type createProductUseCase struct {
	repo repository.ProductRepository
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product) error {
	return c.repo.Add(newProduct)
}

func NewCreateProductUseCase(repo repository.ProductRepository) CreateProductUseCase {
	return &createProductUseCase{
		repo: repo,
	}
}
