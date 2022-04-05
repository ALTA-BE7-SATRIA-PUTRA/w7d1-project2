package project

import (
	_entities "project2/entities"
)

type ProjectUseCaseInterface interface {
	PostProject(project _entities.Project, idToken int) (_entities.Project, int, error)
	GetAll() ([]_entities.Project, error)
	PutProject(id int, project _entities.Project) (_entities.Project, int, error)
	DeleteProject(id int) (_entities.Project, int, error)
}
