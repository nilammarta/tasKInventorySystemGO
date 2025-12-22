package service

import (
	"TaskInventorySystem/models"
	"TaskInventorySystem/repository"
)

type ItemCategoryService struct {
	repo repository.ItemCategoryRepository
}

func NewCategoryService(repo repository.ItemCategoryRepository) *ItemCategoryService {
	return &ItemCategoryService{repository: repo}
}

func (service *ItemCategoryService) GetAllItemCategories() ([]models.ItemCategory, error) {
	return service.repo.FindAll()
}
