package service

import (
	"TaskInventorySystem/models"
	"TaskInventorySystem/repository"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ItemCategoryService struct {
	repo repository.ItemCategoryRepository
}

type ItemService struct {
	repoItem repository.ItemRepository
}

func NewCategoryService(repo repository.ItemCategoryRepository) *ItemCategoryService {
	return &ItemCategoryService{repo: repo}
}

func NewItemService(repoItem repository.ItemRepository) *ItemService {
	return &ItemService{repoItem: repoItem}
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

//func (service *ItemCategoryService) DeleteItemCategory(numberEdit int) error {
//categories, _ := service.repo.FindAll()
//catgory, _ := service.repo.FindByID(categories[numberEdit-1].ID)

// Validasi apakah category sudah terpakai pada data item, jika sudah, tampilkan pesan jangan dihapus
//items, _ := service.repoItem.FindAll()
//}

/*
MENU ITEMS
*/
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
	itemCategory *ItemCategoryService,
	name string,
	stock int,
	categoriesInput []int) (*models.Item, error) {

	// dapatkan data categories berdasarkan id inputan
	categories, _ := itemCategory.repo.FindAll()
	var itemCategories []models.ItemCategory
	for _, input := range categoriesInput {
		id := categories[input-1].ID
		category, _ := itemCategory.repo.FindByID(id)
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

/*
HELPER FUNCTION
*/
func (service *ItemCategoryService) ValidateStringInput(inputString string) (string, error) {
	// remove newline di akhir
	inputString = strings.TrimSpace(inputString)

	if inputString == "" {
		return "", errors.New("Please input the correct task")
	} else {
		return inputString, nil
	}
}

func (service *ItemCategoryService) ValidateIntInput(input string, lengthData int) (int, error) {
	input = strings.TrimSpace(input)
	// convert number string input menjadi int
	numberInput, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Please input a number")
	}

	if numberInput <= lengthData && numberInput >= 1 {
		return numberInput, nil
	} else {
		return -1, errors.New("Please choose the correct task")
	}
}

func ValidateInputAsString(inputString string) (string, error) {
	// remove newline di akhir
	inputString = strings.TrimSpace(inputString)

	if inputString == "" {
		return "", errors.New("Please input the correct task")
	} else {
		return inputString, nil
	}
}

func ValidateInputAsInt(input string) (int, error) {
	input = strings.TrimSpace(input)
	// convert number string input menjadi int
	numberInput, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Please input a number")
	} else {
		return numberInput, nil
	}
}

func ValidateIntInputWithLength(input string, lengthData int) (int, error) {
	input = strings.TrimSpace(input)
	// convert number string input menjadi int
	numberInput, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Please input a number")
	}

	if numberInput <= lengthData && numberInput >= 1 {
		return numberInput, nil
	} else {
		return -1, errors.New("Please choose the correct task")
	}
}

func AskSelectCategoryInput(ask string) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(ask)
	input, _ := reader.ReadString('\n')

	// Hapus newline
	input = strings.TrimSpace(input)

	// Pisahkan berdasarkan koma
	parts := strings.Split(input, ",")

	var numbers []int

	for _, part := range parts {
		// Hilangkan spasi
		part = strings.TrimSpace(part)

		// Konversi ke int
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.New("Input tidak valid!")
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func CheckCategoryLength(itemCategory *ItemCategoryService) (int, error) {
	categories, _ := itemCategory.repo.FindAll()
	if len(categories) == 0 || categories == nil {
		return 0, errors.New("Cannot create item data, add category data first!")
	} else {
		return len(categories), nil
	}
}
