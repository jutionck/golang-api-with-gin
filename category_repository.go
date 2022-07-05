package main

type CategoryRepository interface {
	Insert(newCategory *Category) Category
	List() []Category
}

type categoryRepository struct {
	db []Category
}

func (c *categoryRepository) Insert(newCategory *Category) Category {
	c.db = append(c.db, *newCategory)
	return *newCategory
}

func (c *categoryRepository) List() []Category {
	return c.db
}

func NewCategoryRepository() CategoryRepository {
	repo := new(categoryRepository)
	return repo
}
