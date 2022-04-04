package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name" form:"name"`
	Email    string    `gorm:"unique" json:"email" form:"email"`
	Address  string    `json:"address" form:"address"`
	Password string    `json:"password" form:"password"`
	Task     []Task    `gorm:"foreignKey:IdUser;references:ID"`
	Project  []Project `gorm:"foreignKey:IdUser;references:ID"`
}
