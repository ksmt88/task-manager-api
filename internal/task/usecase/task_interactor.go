package usecase

import "github.com/ksmt88/taskManager-api/internal/task/domain"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Add(t domain.Task) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.Store(t)
	return
}

func (interactor *TaskInteractor) Save(id int, t domain.Task) (err error) {
	err = interactor.TaskRepository.Update(id, t)
	return
}

func (interactor *TaskInteractor) Remove(id int) (err error) {
	err = interactor.TaskRepository.Delete(id)
	return
}

func (interactor *TaskInteractor) All() (tasks domain.Tasks, err error) {
	tasks, err = interactor.TaskRepository.FindAll()
	return
}

func (interactor *TaskInteractor) FindById(id int) (t domain.Task, err error) {
	t, err = interactor.TaskRepository.FindById(id)
	return
}
