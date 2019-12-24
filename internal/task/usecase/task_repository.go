package usecase

import "../domain"

type TaskRepository interface {
	Store(domain.Task) (domain.Task, error)
	Update(int, domain.Task) error
	Delete(int) error
	FindById(int) (domain.Task, error)
	FindAll() (domain.Tasks, error)
}
