package task

import (
	_entities "project2/entities"
)

type TaskUseCaseInterface interface {
	PostTask(task _entities.Task, idToken int) (_entities.Task, int, error)
	GetAll() ([]_entities.Task, error)
	PutTask(id int, task _entities.Task) (_entities.Task, int, error)
	DeleteTask(id int) (_entities.Task, int, error)
	PostTaskComplete(idToken int, task _entities.Task) (_entities.Task, int, error)
	PostTaskReOpen(idToken int, task _entities.Task) (_entities.Task, int, error)
}
