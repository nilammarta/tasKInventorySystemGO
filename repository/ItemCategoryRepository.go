package repository

import "TaskInventorySystem/models"

type ItemCategoryRepository interface {
	FindAll() ([]models.ItemCategory, error)
	FindByID(id int) (models.ItemCategory, error)
	Save(category models.ItemCategory) error
	Update(category models.ItemCategory) error
	Delete(id int) error
}
