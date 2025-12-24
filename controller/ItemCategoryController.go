package controller

import (
	"TaskInventorySystem/service"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ItemCategoryController struct {
	service *service.ItemCategoryService
}

type ItemController struct {
	itemService *service.ItemService
}

func NewItemCategoryController(service *service.ItemCategoryService) *ItemCategoryController {
	return &ItemCategoryController{service: service}
}

func NewItemController(itemService *service.ItemService) *ItemController {
	return &ItemController{itemService: itemService}
}

func (controller *ItemCategoryController) ShowItemCategories() {
	categories, err := controller.service.GetAllItemCategories()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
}

func (controller *ItemCategoryController) AddItemCategory() {
	for {
		fmt.Print("Category name: ")
		reader := bufio.NewReader(os.Stdin)
		inputCategory, _ := reader.ReadString('\n')
		inputCategory = strings.TrimSpace(inputCategory)
		_, err := controller.service.AddNewCategoryItem(inputCategory)
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			fmt.Println("New category added successfully!")
			return
		}
	}
}

func (controller *ItemCategoryController) EditItemCategory() {
	categories, err := controller.service.GetAllItemCategories()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("List Category:")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose category item: ")
		input, _ := reader.ReadString('\n')
		validInt, err := controller.service.ValidateIntInput(input, len(categories))
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			for {
				fmt.Print("Input new category name: ")
				inputString, _ := reader.ReadString('\n')
				validString, err := controller.service.ValidateStringInput(inputString)
				if err != nil {
					fmt.Println(err.Error())
					continue
				} else {
					err := controller.service.EditItemCategory(validInt, validString)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Item category updated successfully!")
					}
					return
				}
			}
		}
	}
}

func (controller *ItemCategoryController) DeleteItemCategory() {
	categories, err := controller.service.GetAllItemCategories()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("List Category:")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}

	//reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose category item: ")
		//input, _ := reader.ReadString('\n')
		//validInt, err := controller.service.ValidateIntInput(input, len(categories))
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			// on kan reading input diatas
			// bagian ini panggil function delete di service
			//di service lakukan validasi apabila data categoty sudah ada di item, maka tidak bisa dihapus.
			//	lakukan delete
		}
	}
}

/*
ITEMS CONTROLLER
*/
func (itemController *ItemController) ShowItems() {
	items, err := itemController.itemService.GetAllItems()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, item := range items {
		fmt.Printf("%d. %s - Stock: %d - Categories: ", i+1, item.Name, item.Stock)

		// cek di go cara menampilkan akar menjadi koma (kyknya cara nay ada)
		for j, category := range item.Categories {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Print(category.Name)
		}
		fmt.Println()
	}
}

func (itemController *ItemController) AddItem(categoryController ItemCategoryController) {

	_, err := service.CheckCategoryLength(categoryController.service)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Item name: ")
		inputString, _ := reader.ReadString('\n')
		name, err := service.ValidateInputAsString(inputString)
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			for {
				fmt.Print("Stock: ")
				inputInt, _ := reader.ReadString('\n')
				stock, err := service.ValidateInputAsInt(inputInt)
				if err != nil {
					fmt.Println(err.Error())
					continue
				} else {
					fmt.Println("Select the number of item category:")
					categoryController.ShowItemCategories()
					for {
						categorySelected, err := service.AskSelectCategoryInput(
							"Example input for multiple categories (1,2,3): ")
						if err != nil {
							fmt.Println(err.Error())
						} else {
							_, err := itemController.itemService.AddNewItem(categoryController.service, name, stock, categorySelected)
							if err != nil {
								fmt.Println(err.Error())
								return
							} else {
								fmt.Println("Item category added successfully!")
								return
							}
						}
					}
				}
			}
		}
	}
}
