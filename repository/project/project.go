package project

import (
	"fmt"
	_entities "project2/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		database: db,
	}
}

func (pr *ProjectRepository) PostProject(project _entities.Project, idToken int) (_entities.Project, error) {

	project.IdUser = uint(idToken)

	tx := pr.database.Save(&project)
	if tx.Error != nil {
		return project, tx.Error
	}
	return project, nil
}
func (pr *ProjectRepository) GetAll() ([]_entities.Project, error) {
	var projects []_entities.Project
	tx := pr.database.Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return projects, nil
}

func (ur *ProjectRepository) PutProject(project _entities.Project, idToken int) (_entities.Project, error) {

	if project.IdUser != uint(idToken) {
		return _entities.Project{}, fmt.Errorf("not autorized")
	}

	project.IdUser = uint(idToken)

	tx := ur.database.Updates(&project)
	if tx != nil {
		return project, tx.Error
	}
	return project, nil
}

func (ur *ProjectRepository) DeleteProject(id int) (_entities.Project, int, error) {
	var project _entities.Project

	tx := ur.database.Where("ID = ?", id).Delete(&project)
	if tx.Error != nil {
		return project, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return project, 0, nil
	}
	return project, int(tx.RowsAffected), nil
}
