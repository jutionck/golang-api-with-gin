package main

type CategoryUseCase interface {
	Insert(newCategory *Category) Category
	List() []Category
}

type categoryUseCase struct {
	repo CategoryRepository
}

func (c *categoryUseCase) Insert(newCategory *Category) Category {
	return c.repo.Insert(newCategory)
}

func (c *categoryUseCase) List() []Category {
	return c.repo.List()
}

func NewCategoryUseCase(repo CategoryRepository) CategoryUseCase {
	uc := new(categoryUseCase)
	uc.repo = repo
	return uc
}
