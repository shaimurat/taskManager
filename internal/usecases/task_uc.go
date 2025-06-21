package usecases

import "taskManager/internal/domain"

type TaskUseCase struct {
	repo domain.TaskRepo
}

func NewTaskUseCase(repo domain.TaskRepo) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}
func (u *TaskUseCase) CreateTask(task domain.Task) error {
	return u.repo.Create(task)
}

func (u *TaskUseCase) UpdateTask(task domain.Task) error {
	return u.repo.Update(task)
}
func (u *TaskUseCase) DeleteTask(task domain.Task) error {
	return u.repo.Delete(task)
}
