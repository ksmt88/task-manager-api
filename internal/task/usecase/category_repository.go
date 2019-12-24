package usecase

import "../domain"

type CategoryRepository interface {
	Store(domain.Category) (int, error)
	Update(int, domain.Category) error
	Delete(int) error
	FindById(int) (domain.Category, error)
	FindAll() (domain.Categories, error)
}
