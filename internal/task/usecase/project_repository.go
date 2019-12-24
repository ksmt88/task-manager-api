package usecase

import "../domain"

type ProjectRepository interface {
	Store(domain.Project) (int, error)
	Update(int, domain.Project) error
	Delete(int) error
	FindById(int) (domain.Project, error)
	FindAll() (domain.Projects, error)
}
