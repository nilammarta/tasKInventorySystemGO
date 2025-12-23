package main

import (
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
	fmt.Println("3. EDIT CATEGORY STOCK")
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

func manageCategories(itemCategories *[]ItemCategory) {
	var categoryMenu string

	for {
		showCategoryMenu()
		fmt.Scanf("%s", &categoryMenu)

		if categoryMenu == "1" {
			fmt.Println("===== SHOW CATEGORIES =====")
			showCategories(*itemCategories)
			continue
		} else if categoryMenu == "2" {
			fmt.Println("===== ADD CATEGORY =====")
			continue
		} else if categoryMenu == "3" {
			fmt.Println("===== EDIT CATEGORY =====")
			continue
		} else if categoryMenu == "4" {
			fmt.Println("===== REMOVE CATEGORY =====")
			continue
		} else if categoryMenu == "5" {
			fmt.Println(" <-- BACK TO MAIN MENU")
			break
		}
	}
}

/*
FUNCTION TO MANAGE CATEGORIES
*/
func showCategories(categories []ItemCategory) {
	if categories == nil || len(categories) == 0 {
		fmt.Println("Category item is empty!")
	} else {
		for i, category := range categories {
			fmt.Printf("%d. %s\n", i+1, category.Name)
		}
	}
	pressEnterToContinue()
}

// add categories
func addCategory(category *[]ItemCategory) {

}

func main() {

	/*
		REPOSITORY & SERVICE
	*/
	categoryRepo := repository.NewInMemoryItemCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo)

	service := service.ItemCategoryService{repo: repo}

	var mainMenu string

	// inisiasi slice
	var itemCategories []ItemCategory
	//var items []Item

	for {
		showMainMenu()
		fmt.Scanf("%s", &mainMenu)

		if mainMenu == "1" {
			fmt.Println("===== SHOW ITEMS =====")

			continue
		} else if mainMenu == "2" {
			fmt.Println("===== ADD ITEM =====")
			continue
		} else if mainMenu == "3" {
			fmt.Println("===== UPDATE ITEM STOCK =====")
			continue
		} else if mainMenu == "4" {
			fmt.Println("===== REMOVE ITEM =====")
			continue
		} else if mainMenu == "5" {
			// panggil function yg menangani manu ini
			manageCategories(&itemCategories)
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
