package repository

import "TaskInventorySystem/models"

type ItemRepository interface {
	FindAll() ([]models.Item, error)
	FindByID(id int) (models.Item, error)
	Save(item models.Item) error
	Update(item models.Item) error
	Delete(id int) error
}
