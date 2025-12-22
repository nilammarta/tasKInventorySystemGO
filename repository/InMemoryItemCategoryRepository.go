package repository

import (
	"TaskInventorySystem/models"
	"errors"
)

type InMemoryItemCategoryRepository struct {
	categories []models.ItemCategory
}

// Buat repository
func NewInMemoryItemCategoryRepository() ItemCategoryRepository {
	return &InMemoryItemCategoryRepository{
		categories: []models.ItemCategory{},
	}
}

func (repository *InMemoryItemCategoryRepository) FindAll() ([]models.ItemCategory, error) {
	return repository.categories, nil
}

func (repository *InMemoryItemCategoryRepository) FindByID(id int) (models.ItemCategory, error) {
	for _, category := range repository.categories {
		if category.ID == id {
			return category, nil
		}
	}
	return models.ItemCategory{}, errors.New("Item category not found!")
}

func (repository *InMemoryItemCategoryRepository) Save(newCategory models.ItemCategory) error {
	repository.categories = append(repository.categories, newCategory)
	return nil
}

func (repository *InMemoryItemCategoryRepository) Update(updateCategory models.ItemCategory) error {
	for i, category := range repository.categories {
		if category.ID == updateCategory.ID {
			repository.categories[i] = updateCategory
			return nil
		}
	}
	return errors.New("Item category not found!")
}

func (repository *InMemoryItemCategoryRepository) Delete(id int) error {
	for i, category := range repository.categories {
		if category.ID == id {
			repository.categories = append(repository.categories[:i], repository.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("Item category not found!")
}
