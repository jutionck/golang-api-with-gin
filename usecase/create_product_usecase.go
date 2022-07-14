package usecase

import (
	"fmt"
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
	"mime/multipart"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error
}

type createProductUseCase struct {
	repo     repository.ProductRepository
	fileRepo repository.FileRepository
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error {
	fileName := fmt.Sprintf("img-%s.%s", newProduct.ProductId, fileExt)
	fileLocation, err := c.fileRepo.Save(file, fileName)
	if err != nil {
		return err
	}
	newProduct.ImgPath = fileLocation
	err = c.repo.Add(newProduct)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateProductUseCase(repo repository.ProductRepository, fileRepo repository.FileRepository) CreateProductUseCase {
	return &createProductUseCase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
