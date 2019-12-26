package usecase

import "github.com/ksmt88/task-manager-api/internal/task/domain"

type ProjectRepository interface {
	Store(domain.Project) (int, error)
	Update(int, domain.Project) error
	Delete(int) error
	FindById(int) (domain.Project, error)
	FindAll() (domain.Projects, error)
}
