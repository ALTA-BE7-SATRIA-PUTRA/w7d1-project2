package task

import (
	"fmt"
	_entities "project2/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (pr *TaskRepository) PostTask(task _entities.Task, idToken int) (_entities.Task, error) {

	task.IdUser = uint(idToken)

	tx := pr.database.Save(&task)
	if tx.Error != nil {
		return task, tx.Error
	}
	return task, nil
}
func (pr *TaskRepository) GetAll() ([]_entities.Task, error) {
	var tasks []_entities.Task
	tx := pr.database.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}

func (ur *TaskRepository) PutTask(task _entities.Task, idToken int) (_entities.Task, error) {

	if task.IdUser != uint(idToken) {
		return _entities.Task{}, fmt.Errorf("not autorized")
	}

	task.IdUser = uint(idToken)

	tx := ur.database.Updates(&task)
	if tx != nil {
		return task, tx.Error
	}
	return task, nil
}

func (ur *TaskRepository) DeleteTask(id int) (_entities.Task, int, error) {
	var task _entities.Task

	tx := ur.database.Where("ID = ?", id).Delete(&task)
	if tx.Error != nil {
		return task, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, 0, nil
	}
	return task, int(tx.RowsAffected), nil
}
func (ur *TaskRepository) PostTaskComplete(task _entities.Task, idToken int) (_entities.Task, error) {

	if task.IdUser != uint(idToken) {
		return _entities.Task{}, fmt.Errorf("not autorized")
	}

	task.IdUser = uint(idToken)
	task.Status = "Completed"
	tx := ur.database.Updates(&task)
	if tx != nil {
		return task, tx.Error
	}
	return task, nil
}
func (ur *TaskRepository) PostTaskReOpen(task _entities.Task, idToken int) (_entities.Task, error) {

	if task.IdUser != uint(idToken) {
		return _entities.Task{}, fmt.Errorf("not autorized")
	}

	task.IdUser = uint(idToken)
	task.Status = "Re Open"
	tx := ur.database.Updates(&task)
	if tx != nil {
		return task, tx.Error
	}
	return task, nil
}
