package repository

import "github.com/jutionck/golang-api-with-gin/model"

type ProductRepository interface {
	Add(newpProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

type productRepository struct {
	productDb *[]model.Product
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	return *p.productDb, nil
}

func (p *productRepository) Add(newpProduct *model.Product) error {
	*p.productDb = append(*p.productDb, *newpProduct)
	return nil
}

func NewProductRepository(productDb *[]model.Product) ProductRepository {
	repo := new(productRepository)
	repo.productDb = productDb
	return repo
}
