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

func NewItemCategoryController(service *service.ItemCategoryService) *ItemCategoryController {
	return &ItemCategoryController{service: service}
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
				confirmInput, err := service.AskInputAsString("Are you sure you want to delete category item (y/n): ")
				if err != nil {
					fmt.Println(err.Error())
				}
				if confirmInput == "y" || confirmInput == "Y" {
					err := controller.service.DeleteItemCategory(validInt)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Category data deleted successfully!")
					}
					return
				} else if confirmInput == "n" || confirmInput == "N" {
					fmt.Println("Category deletion has been cancelled!")
					return
				} else {
					fmt.Println("Please input y/n")
					continue
				}
			}
		}
	}
}
