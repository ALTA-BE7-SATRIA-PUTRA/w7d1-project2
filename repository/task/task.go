package task

import (
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
