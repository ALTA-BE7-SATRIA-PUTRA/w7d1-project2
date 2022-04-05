package project

import (
	"net/http"
	"project2/delivery/helper"
	_middlewares "project2/delivery/middlewares"
	_entities "project2/entities"
	_projectUseCase "project2/usecase/project"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectUseCase _projectUseCase.ProjectUseCaseInterface
}

func NewProjectHandler(projectUseCase _projectUseCase.ProjectUseCaseInterface) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}
func (ph *ProjectHandler) PostProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var project _entities.Project
		tx := c.Bind(&project)

		projects, _ := ph.projectUseCase.GetAll()
		for _, val := range projects {
			if val.Project == project.Project {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("project already exist"))
			}
		}
		idToken, _ := _middlewares.ExtractToken(c)
		newProject, _, err := ph.projectUseCase.PostProject(project, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create project"))
		}
		if tx != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create new project"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to insert new data", newProject))

	}
}
func (ph *ProjectHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, err := ph.projectUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get all data", projects))
	}
}

func (ph *ProjectHandler) PutProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		var project _entities.Project
		var updateProject _entities.Project

		c.Bind(&updateProject)
		if updateProject.IdUser == 0 {
			project.IdUser = uint(idToken)
		}
		if updateProject.Project != "" {
			project.Project = updateProject.Project
		}

		if updateProject.Description != "" {
			project.Description = updateProject.Description
		}

		project, _, err := ph.projectUseCase.PutProject(idToken, project)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to update project"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to update Project", project))
	}
}

func (ph *ProjectHandler) DeleteProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, errorconv := strconv.Atoi(idStr)

		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be int")
		}

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		projects, rows, err := ph.projectUseCase.DeleteProject(id)
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete Project"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to delete Project", projects))
	}
}
