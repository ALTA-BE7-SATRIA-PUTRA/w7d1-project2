package task

import (
	_taskUseCase "project2/usecase/task"
)

type TaskHandler struct {
	taskUseCase _taskUseCase.TaskUseCaseInterface
}

func NewTaskHandler(taskUseCase _taskUseCase.TaskUseCaseInterface) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}
