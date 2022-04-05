package task

import (
	_entities "project2/entities"
)

type TaskRepositoryInterface interface {
	PostTask(task _entities.Task, idToken int) (_entities.Task, error)
	GetAll() ([]_entities.Task, error)
	PutTask(project _entities.Task, id int) (_entities.Task, error)
	DeleteTask(id int) (_entities.Task, int, error)
	PostTaskComplete(project _entities.Task, id int) (_entities.Task, error)
	PostTaskReOpen(project _entities.Task, id int) (_entities.Task, error)
}
