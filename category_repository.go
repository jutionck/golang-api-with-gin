package main

type CategoryRepository interface {
	Insert(newCategory *Category) Category
	List() []Category
}

type categoryRepository struct {
	db []Category
}

func (c *categoryRepository) Insert(newCategory *Category) Category {
	var category Category
	category = Category{
		Id:   newCategory.Id,
		Name: newCategory.Name,
	}
	c.db = append(c.db, category)
	return category
}

func (c *categoryRepository) List() []Category {
	var categories []Category
	for _, category := range c.db {
		categories = append(categories, category)
	}
	return categories
}

func NewCategoryRepository(db []Category) CategoryRepository {
	repo := new(categoryRepository)
	repo.db = db
	return repo
}
