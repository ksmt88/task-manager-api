package usecase

import "github.com/ksmt88/task-manager-api/internal/task/domain"

type ProjectInteractor struct {
	ProjectRepository ProjectRepository
}

func (interactor *ProjectInteractor) Add(p domain.Project) (id int, err error) {
	id, err = interactor.ProjectRepository.Store(p)
	return
}

func (interactor *ProjectInteractor) Save(id int, p domain.Project) (err error) {
	err = interactor.ProjectRepository.Update(id, p)
	return
}

func (interactor *ProjectInteractor) Remove(id int) (err error) {
	err = interactor.ProjectRepository.Delete(id)
	return
}

func (interactor *ProjectInteractor) All() (projects domain.Projects, err error) {
	projects, err = interactor.ProjectRepository.FindAll()
	return
}

func (interactor *ProjectInteractor) FindById(id int) (p domain.Project, err error) {
	p, err = interactor.ProjectRepository.FindById(id)
	return
}
