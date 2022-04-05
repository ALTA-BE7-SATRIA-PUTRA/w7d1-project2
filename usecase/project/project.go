package project

import (
	_entities "project2/entities"
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
func (puc *ProjectUseCase) PostProject(project _entities.Project, idToken int) (_entities.Project, int, error) {
	projects, err := puc.projectRepository.PostProject(project, idToken)
	return projects, idToken, err
}

func (puc *ProjectUseCase) GetAll() ([]_entities.Project, error) {
	projects, err := puc.projectRepository.GetAll()
	return projects, err
}

func (uuc *ProjectUseCase) PutProject(id int, project _entities.Project) (_entities.Project, int, error) {
	project.ID = uint(id)

	projectPut, errPut := uuc.projectRepository.PutProject(project, id)
	return projectPut, id, errPut
}

func (uuc *ProjectUseCase) DeleteProject(id int) (_entities.Project, int, error) {
	project, rows, err := uuc.projectRepository.DeleteProject(id)
	if project.Project == "" {
		return project, rows, err
	}
	return project, rows, nil
}
