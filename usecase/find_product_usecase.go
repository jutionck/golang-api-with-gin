package usecase

import (
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
)

type FindProductUseCase interface {
	ById(id string) (model.Product, error)
}

type findProductUseCase struct {
	repo repository.ProductRepository
}

func (l *findProductUseCase) ById(id string) (model.Product, error) {
	return l.repo.FindById(id)
}

func NewFindProductUseCase(repo repository.ProductRepository) FindProductUseCase {
	return &findProductUseCase{repo: repo}
}
