package project

import (
	_projectUseCase "project2/usecase/project"
)

type ProjectHandler struct {
	projectUseCase _projectUseCase.ProjectUseCaseInterface
}

func NewProjectHandler(projectUseCase _projectUseCase.ProjectUseCaseInterface) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}
