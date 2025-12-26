package service

import (
	"TaskInventorySystem/models"
	"TaskInventorySystem/repository"
	"errors"
)

type ItemCategoryService struct {
	repo     repository.ItemCategoryRepository
	itemRepo repository.ItemRepository
}

func NewCategoryService(
	repo repository.ItemCategoryRepository,
	itemRepo repository.ItemRepository,
) *ItemCategoryService {
	return &ItemCategoryService{
		repo:     repo,
		itemRepo: itemRepo,
	}
}

func (service *ItemCategoryService) GetAllItemCategories() ([]models.ItemCategory, error) {
	categories, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if categories == nil || len(categories) == 0 {
		return nil, errors.New("Item category data is empty!")
	} else {
		return categories, nil
	}
}

func (service *ItemCategoryService) AddNewCategoryItem(categoryName string) (*models.ItemCategory, error) {
	if categoryName == "" {
		return nil, errors.New("Item category could not be empty, please input the correct name!")
	} else {
		categories, _ := service.repo.FindAll()
		var id int
		if categories == nil || len(categories) == 0 {
			id = 1
		} else {
			id = categories[len(categories)-1].ID + 1
		}

		// buat data dengan tipe Item Category
		newCategory := models.ItemCategory{
			ID:   id,
			Name: categoryName,
		}

		// save data
		err := service.repo.Save(newCategory)
		if err != nil {
			return nil, err
		} else {
			return &newCategory, nil
		}
	}
}

func (service *ItemCategoryService) EditItemCategory(numberEdit int, newCategory string) error {
	categories, _ := service.repo.FindAll()
	categoryEdit := models.ItemCategory{
		ID:   categories[numberEdit-1].ID,
		Name: newCategory,
	}
	return service.repo.Update(categoryEdit)
}

func (service *ItemCategoryService) DeleteItemCategory(numberInput int) error {
	categories, _ := service.repo.FindAll()
	categoryID := categories[numberInput-1].ID
	// Validasi apakah category sudah terpakai pada data item, jika sudah, tampilkan pesan jangan dihapus
	items, err := service.itemRepo.FindAll()
	if err != nil {
		return err
	}

	// validasi apakah data kategori telah ditambahkan di data item
	for _, item := range items {
		for _, category := range item.Categories {
			if category.ID == categoryID {
				return errors.New("Item category is in use cannot be deleted!")
			}
		}
	}

	return service.repo.Delete(categoryID)
}
