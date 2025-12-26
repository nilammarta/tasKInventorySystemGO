package repository

import (
	"TaskInventorySystem/models"
	"errors"
)

type InMemoryItemRepository struct {
	items []models.Item
}

func NewInMemoryItemRepository() ItemRepository {
	return &InMemoryItemRepository{
		items: []models.Item{},
	}
}
func (repository *InMemoryItemRepository) FindAll() ([]models.Item, error) {

	return repository.items, nil
}

func (repository *InMemoryItemRepository) FindByID(id int) (models.Item, error) {
	for _, item := range repository.items {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item not found!")
}

func (repository *InMemoryItemRepository) Save(newItem models.Item) error {
	repository.items = append(repository.items, newItem)
	return nil
}
func (repository *InMemoryItemRepository) Update(updateItem models.Item) error {
	for i, item := range repository.items {
		if item.ID == updateItem.ID {
			repository.items[i] = updateItem
			return nil
		}
	}
	return errors.New("Item not found!")
}
func (repository *InMemoryItemRepository) Delete(id int) error {
	for i, item := range repository.items {
		if item.ID == id {
			repository.items = append(repository.items[:i], repository.items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item category not found!")
}
