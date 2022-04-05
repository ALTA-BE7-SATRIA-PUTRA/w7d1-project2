package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Project     string `json:"project" form:"project"`
	Description string `json:"description" form:"description"`
	IdUser      uint   `json:"id_user" form:"id_user"`
	Task        []Task `gorm:"foreignKey:IdUser;preferences:ID"`
}
