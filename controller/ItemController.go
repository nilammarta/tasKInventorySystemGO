package controller

import (
	"TaskInventorySystem/service"
	"bufio"
	"fmt"
	"os"
)

/*
ITEMS CONTROLLER
*/
type ItemController struct {
	itemService     *service.ItemService
	categoryService *service.ItemCategoryService
}

func NewItemController(
	itemService *service.ItemService,
	categoryService service.ItemCategoryService,
) *ItemController {
	return &ItemController{
		itemService:     itemService,
		categoryService: &categoryService,
	}
}

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

func (itemController *ItemController) AddItem() {

	categories, _ := itemController.categoryService.GetAllItemCategories()
	err := service.CheckCategoryData(categories)
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
					for i, category := range categories {
						fmt.Printf("%d. %s\n", i+1, category.Name)
					}
					for {
						categorySelected, err := service.AskSelectCategoryInput(
							"Example input for multiple categories (1,2,3): ")
						if err != nil {
							fmt.Println(err.Error())
						} else {
							_, err := itemController.itemService.AddNewItem(name, stock, categorySelected)
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

func (itemController *ItemController) EditItemStock() {
	items, _ := itemController.itemService.GetAllItems()
	if len(items) == 0 {
		fmt.Println("Items data is empty!")
		return
	}

	for {
		// Ask select item
		numberInput, err := service.AskInputAsInt("Your Choices: ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Input the the new stock:")
			for {
				newStock, err := service.AskInputAsInt("Stock: ")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					err := itemController.itemService.EditItemStock(numberInput, newStock)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Item stock updated successfully!")
						return
					}
				}
			}
		}
	}
}

func (itemController *ItemController) DeleteItem() {
	items, _ := itemController.itemService.GetAllItems()
	if len(items) == 0 {
		fmt.Println("Items data is empty!")
		return
	}

	fmt.Println("Select item to update their stock:")
	itemController.ShowItems()
	for {
		// Ask select item
		numberInput, err := service.AskInputIntInRange("Your Choices: ", len(items))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for {
				message, err := itemController.itemService.DeleteItem(numberInput)
				if err != nil {
					fmt.Println(err.Error())
					return
				} else if message != "" {
					fmt.Println(message)
					continue
				} else {
					fmt.Println("Item deleted successfully!")
					return
				}
			}
		}
	}
}
