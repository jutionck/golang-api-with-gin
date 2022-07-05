package manager

import "github.com/jutionck/golang-api-with-gin/usecase"

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUseCase
	ListProductUseCase() usecase.ListProductUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ListProductUseCase() usecase.ListProductUseCase {
	return usecase.NewListProductUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
