package usecase

import "github.com/ksmt88/taskManager-api/internal/task/domain"

type CategoryInteractor struct {
	CategoryRepository CategoryRepository
}

func (interactor *CategoryInteractor) Add(c domain.Category) (id int, err error) {
	id, err = interactor.CategoryRepository.Store(c)
	return
}

func (interactor *CategoryInteractor) Save(id int, c domain.Category) (err error) {
	err = interactor.CategoryRepository.Update(id, c)
	return
}

func (interactor *CategoryInteractor) Remove(id int) (err error) {
	err = interactor.CategoryRepository.Delete(id)
	return
}

func (interactor *CategoryInteractor) All() (categories domain.Categories, err error) {
	categories, err = interactor.CategoryRepository.FindAll()
	return
}

func (interactor *CategoryInteractor) FindById(id int) (c domain.Category, err error) {
	c, err = interactor.CategoryRepository.FindById(id)
	return
}
