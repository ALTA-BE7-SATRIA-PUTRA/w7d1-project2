package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
	Status      string `gorm:"default:Not Completed" json:"status" from:"status"`
	IdUser      uint   `json:"id_user" form:"id_user"`
	IdProject   uint   `json:"id_project" form:"id_project"`
}
