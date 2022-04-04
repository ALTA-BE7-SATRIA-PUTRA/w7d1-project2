package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task        string `json:"task" form:"task"`
	Description string `json:"description" form:"description"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProject   string `json:"id_project" form:"id_project"`
}
