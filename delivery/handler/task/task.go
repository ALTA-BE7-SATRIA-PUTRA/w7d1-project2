package task

import (
	"net/http"
	"project2/delivery/helper"
	_middlewares "project2/delivery/middlewares"
	_entities "project2/entities"
	_taskUseCase "project2/usecase/task"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase _taskUseCase.TaskUseCaseInterface
}

func NewTaskHandler(taskUseCase _taskUseCase.TaskUseCaseInterface) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}
func (ph *TaskHandler) PostTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var task _entities.Task
		tx := c.Bind(&task)
		tasks, _ := ph.taskUseCase.GetAll()
		for _, val := range tasks {
			if val.Task == task.Task {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("task already exist"))
			}
		}
		idToken, _ := _middlewares.ExtractToken(c)
		newTask, _, err := ph.taskUseCase.PostTask(task, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create task"))
		}
		if tx != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create new task"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to insert new data", newTask))

	}
}
func (ph *TaskHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks, err := ph.taskUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get all data", tasks))
	}
}

func (ph *TaskHandler) PutTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		var task _entities.Task
		var updateTask _entities.Task

		c.Bind(&updateTask)

		if updateTask.IdUser == 0 {
			task.IdUser = uint(idToken)
		}
		if updateTask.Task != "" {
			task.Task = updateTask.Task
		}

		if updateTask.Description != "" {
			task.Description = updateTask.Description
		}

		task, _, err := ph.taskUseCase.PutTask(idToken, task)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to update task"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to update Task", task))
	}
}

func (ph *TaskHandler) DeleteTaskHandler() echo.HandlerFunc {
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

		tasks, rows, err := ph.taskUseCase.DeleteTask(id)
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete Task"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses to delete Task", tasks))
	}
}
func (ph *TaskHandler) PostTaskCompleteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		var task _entities.Task
		var updateTask _entities.Task

		c.Bind(&updateTask)

		if updateTask.IdUser == 0 {
			task.IdUser = uint(idToken)
		}
		if updateTask.Task != "" {
			task.Task = updateTask.Task
		}

		if updateTask.Description != "" {
			task.Description = updateTask.Description
		}

		task, _, err := ph.taskUseCase.PostTaskComplete(idToken, task)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to completed task"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Task Completed", task))
	}
}
func (ph *TaskHandler) PostTaskReOpenHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		var task _entities.Task
		var updateTask _entities.Task

		c.Bind(&updateTask)

		if updateTask.IdUser == 0 {
			task.IdUser = uint(idToken)
		}
		if updateTask.Task != "" {
			task.Task = updateTask.Task
		}

		if updateTask.Description != "" {
			task.Description = updateTask.Description
		}

		task, _, err := ph.taskUseCase.PostTaskReOpen(idToken, task)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to Re Open task"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Task Opened", task))
	}
}
