package usecase

import "github.com/ksmt88/taskManager-api/internal/task/domain"

type TaskRepository interface {
	Store(domain.Task) (domain.Task, error)
	Update(int, domain.Task) error
	Delete(int) error
	FindById(int) (domain.Task, error)
	FindAll() (domain.Tasks, error)
}
