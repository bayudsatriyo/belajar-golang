package repository

import "golang-unit-test/entity"

type CategoryReository interface {
	FindById(id string) *entity.Category
}
