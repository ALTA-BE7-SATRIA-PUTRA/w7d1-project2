package task

import (
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
