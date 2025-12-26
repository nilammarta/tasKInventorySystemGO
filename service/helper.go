package service

import (
	"TaskInventorySystem/models"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func AskInputAsString(ask string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ask)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("Input is invalid, please input correctly!")
	} else {
		return input, nil
	}
}

func AskInputAsInt(ask string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ask)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// convert number string input menjadi int
	numberInput, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Please input a number")
	} else {
		return numberInput, nil
	}
}

func AskInputIntInRange(ask string, max int) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ask)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// convert number string input menjadi int
	numberInput, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Please input a number")
	}

	if numberInput <= max && numberInput >= 1 {
		return numberInput, nil
	} else {
		return -1, errors.New("Please choose the correct number")
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

		if num > 0 && num < len(parts) {
			numbers = append(numbers, num)
		} else {
			return nil, errors.New("Input tidak valid!")
		}
	}
	return numbers, nil
}

func CheckCategoryData(categories []models.ItemCategory) error {
	if len(categories) == 0 || categories == nil {
		return errors.New("Cannot create item data, add category data first!")
	} else {
		return nil
	}
}
