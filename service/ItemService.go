package service

import (
	"TaskInventorySystem/models"
	"TaskInventorySystem/repository"
	"errors"
	"fmt"
)

type ItemService struct {
	repoItem     repository.ItemRepository
	repoCategory repository.ItemCategoryRepository
}

func NewItemService(
	repoItem repository.ItemRepository,
	repoCategory repository.ItemCategoryRepository,
) *ItemService {
	return &ItemService{
		repoItem:     repoItem,
		repoCategory: repoCategory,
	}
}

func (service *ItemService) GetAllItems() ([]models.Item, error) {
	items, err := service.repoItem.FindAll()
	if err != nil {
		return nil, err
	}

	if items == nil || len(items) == 0 {
		return nil, errors.New("Item data is empty!")
	} else {
		return items, nil
	}
}

func (service *ItemService) AddNewItem(
	name string,
	stock int,
	categoriesInput []int) (*models.Item, error) {

	// dapatkan data categories berdasarkan id inputan
	categories, _ := service.repoCategory.FindAll()
	var itemCategories []models.ItemCategory
	for _, input := range categoriesInput {
		id := categories[input-1].ID
		category, _ := service.repoCategory.FindByID(id)
		itemCategories = append(itemCategories, category)
	}

	// buatkan data item baru:
	newItem := models.Item{
		Name:       name,
		Stock:      stock,
		Categories: itemCategories,
	}

	// save data
	err := service.repoItem.Save(newItem)
	if err != nil {
		return nil, err
	} else {
		return &newItem, nil
	}
}

func (service *ItemService) EditItemStock(numberInput int, newStock int) error {
	items, err := service.repoItem.FindAll()
	if err != nil {
		return err
	}

	if numberInput > len(items) || numberInput < 1 {
		return errors.New("Number of items is out of range!")
	}

	if newStock < 0 {
		return errors.New("Number of stock could not be negative!")
	}

	itemEdit := models.Item{
		ID:         items[numberInput-1].ID,
		Name:       items[numberInput-1].Name,
		Stock:      newStock,
		Categories: items[numberInput-1].Categories,
	}
	return service.repoItem.Update(itemEdit)
}

func (service *ItemService) DeleteItem(numberInput int) (string, error) {
	// Ask comfirm
	confirmInput, err := AskInputAsString("Are you sure you want to delete category item (y/n): ")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	items, err := service.repoItem.FindAll()
	if err != nil {
		return "", err
	}

	if confirmInput == "y" || confirmInput == "Y" {
		itemID := items[numberInput-1].ID
		return "", service.repoItem.Delete(itemID)
	} else if confirmInput == "n" || confirmInput == "N" {
		return "", errors.New("Item deletion has been cancelled!")
	} else {
		return "Invalid confirm input", nil
	}
}
