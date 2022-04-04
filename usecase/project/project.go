package project

import (
	_projectRepository "project2/repository/project"
)

type ProjectUseCase struct {
	projectRepository _projectRepository.ProjectRepositoryInterface
}

func NewProjectUseCase(projectRepo _projectRepository.ProjectRepositoryInterface) ProjectUseCaseInterface {
	return &ProjectUseCase{
		projectRepository: projectRepo,
	}
}
