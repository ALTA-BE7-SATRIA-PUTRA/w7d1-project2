package task

import (
	_entities "project2/entities"
	_taskRepository "project2/repository/task"
)

type TaskUseCase struct {
	taskRepository _taskRepository.TaskRepositoryInterface
}

func NewTaskUseCase(taskRepo _taskRepository.TaskRepositoryInterface) TaskUseCaseInterface {
	return &TaskUseCase{
		taskRepository: taskRepo,
	}
}
func (puc *TaskUseCase) PostTask(task _entities.Task, idToken int) (_entities.Task, int, error) {
	tasks, err := puc.taskRepository.PostTask(task, idToken)
	return tasks, idToken, err
}

func (puc *TaskUseCase) GetAll() ([]_entities.Task, error) {
	tasks, err := puc.taskRepository.GetAll()
	return tasks, err
}

func (uuc *TaskUseCase) PutTask(id int, task _entities.Task) (_entities.Task, int, error) {
	task.ID = uint(id)

	taskPut, errPut := uuc.taskRepository.PutTask(task, id)
	return taskPut, id, errPut
}

func (uuc *TaskUseCase) DeleteTask(id int) (_entities.Task, int, error) {
	task, rows, err := uuc.taskRepository.DeleteTask(id)
	if task.Task == "" {
		return task, rows, err
	}
	return task, rows, nil
}
func (uuc *TaskUseCase) PostTaskComplete(id int, task _entities.Task) (_entities.Task, int, error) {
	task.ID = uint(id)

	taskPut, errPut := uuc.taskRepository.PostTaskComplete(task, id)
	return taskPut, id, errPut
}
func (uuc *TaskUseCase) PostTaskReOpen(id int, task _entities.Task) (_entities.Task, int, error) {
	task.ID = uint(id)

	taskPut, errPut := uuc.taskRepository.PostTaskReOpen(task, id)
	return taskPut, id, errPut
}
