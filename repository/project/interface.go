package project

import (
	_entities "project2/entities"
)

type ProjectRepositoryInterface interface {
	PostProject(project _entities.Project, idToken int) (_entities.Project, error)
	GetAll() ([]_entities.Project, error)
	PutProject(project _entities.Project, id int) (_entities.Project, error)
	DeleteProject(id int) (_entities.Project, int, error)
}
