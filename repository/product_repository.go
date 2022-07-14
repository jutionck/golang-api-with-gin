package repository

import (
	"github.com/jutionck/golang-api-with-gin/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newpProduct *model.Product) error
	Retrieve() ([]model.Product, error)
	FindById(id string) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) FindById(id string) (model.Product, error) {
	var product model.Product
	err := p.db.First(&product, "product_id=?", id).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
func (p *productRepository) Add(newpProduct *model.Product) error {
	err := p.db.Create(&newpProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
