package project

import (
	"gorm.io/gorm"
)

type ProcjectRepository struct {
	database *gorm.DB
}

func NewProcjectRepository(db *gorm.DB) *ProcjectRepository {
	return &ProcjectRepository{
		database: db,
	}
}
