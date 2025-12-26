package main

import (
	"TaskInventorySystem/controller"
	"TaskInventorySystem/repository"
	"TaskInventorySystem/service"
	"fmt"
)

func showMainMenu() {
	fmt.Println("===== MAIN MENU INVENTORY SYSTEM =====")
	fmt.Println("1. SHOW ITEMS")
	fmt.Println("2. ADD ITEM")
	fmt.Println("3. UPDATE ITEM STOCK")
	fmt.Println("4. REMOVE ITEM")
	fmt.Println("5. MANAGE CATEGORIES")
	fmt.Println("6. EXIT")
	fmt.Print("Your choice: ")
}

func showCategoryMenu() {
	fmt.Println("===== CATEGORY MANAGEMENT MENU =====")
	fmt.Println("1. SHOW CATEGORIES")
	fmt.Println("2. ADD CATEGORY")
	fmt.Println("3. EDIT CATEGORY")
	fmt.Println("4. REMOVE CATEGORY")
	fmt.Println("5. BACK TO MAIN MENU")
	fmt.Print("Your choice: ")
}

func pressEnterToContinue() {
	var input string
	fmt.Print("Press ENTER to continue")
	fmt.Scanf("%s", &input)
	fmt.Println(" ")
}

func manageCategories(categoryController *controller.ItemCategoryController) {
	var categoryMenu string

	for {
		showCategoryMenu()
		fmt.Scanf("%s", &categoryMenu)

		if categoryMenu == "1" {
			fmt.Println("===== SHOW CATEGORIES =====")
			categoryController.ShowItemCategories()
			pressEnterToContinue()
			continue
		} else if categoryMenu == "2" {
			fmt.Println("===== ADD CATEGORY =====")
			categoryController.AddItemCategory()
			pressEnterToContinue()
			continue
		} else if categoryMenu == "3" {
			fmt.Println("===== EDIT CATEGORY =====")
			categoryController.EditItemCategory()
			pressEnterToContinue()
			continue
		} else if categoryMenu == "4" {
			fmt.Println("===== REMOVE CATEGORY =====")
			categoryController.DeleteItemCategory()
			pressEnterToContinue()
			continue
		} else if categoryMenu == "5" {
			fmt.Println(" <<< BACK TO MAIN MENU")
			break
		}
	}
}

func main() {
	/*
		REPOSITORY & SERVICE
	*/
	categoryRepo := repository.NewInMemoryItemCategoryRepository()
	itemRepo := repository.NewInMemoryItemRepository()

	categoryService := service.NewCategoryService(categoryRepo, itemRepo)
	itemService := service.NewItemService(itemRepo, categoryRepo)

	categoryController := controller.NewItemCategoryController(categoryService)
	itemController := controller.NewItemController(itemService, *categoryService)

	var mainMenu string

	for {
		showMainMenu()
		fmt.Scanf("%s", &mainMenu)

		if mainMenu == "1" {
			fmt.Println("===== SHOW ITEMS =====")
			itemController.ShowItems()
			pressEnterToContinue()
			continue
		} else if mainMenu == "2" {
			fmt.Println("===== ADD ITEM =====")
			itemController.AddItem()
			pressEnterToContinue()
			continue
		} else if mainMenu == "3" {
			fmt.Println("===== UPDATE ITEM STOCK =====")
			itemController.EditItemStock()
			pressEnterToContinue()
			continue
		} else if mainMenu == "4" {
			fmt.Println("===== REMOVE ITEM =====")
			itemController.DeleteItem()
			pressEnterToContinue()
			continue
		} else if mainMenu == "5" {
			fmt.Println(">>> Go To Category Menu")
			// panggil function yg menangani manu ini
			manageCategories(categoryController)
			continue
		} else if mainMenu == "6" {
			fmt.Println("===== EXIT =====")
			fmt.Print("See you next time!")
			break
		} else {
			fmt.Println("Incorrect menu input!")
			pressEnterToContinue()
		}
	}
}
