package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryReository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	result := service.Repository.FindById(id)
	if result != nil {
		return result, nil
	}

	return result, errors.New("Category not found")
}
