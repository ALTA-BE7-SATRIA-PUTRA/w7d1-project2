package routes

import (
	_authHandler "project2/delivery/handler/auth"
	_projectHandler "project2/delivery/handler/project"
	_taskHandler "project2/delivery/handler/task"
	_userHandler "project2/delivery/handler/user"
	_middlewares "project2/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, ph *_projectHandler.ProjectHandler, th *_taskHandler.TaskHandler) {
	e.POST("/users", uh.PostUserHandler())
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.PutUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())

	e.POST("/projects", ph.PostProjectHandler(), _middlewares.JWTMiddleware())
	e.GET("/Projects", ph.GetAllHandler(), _middlewares.JWTMiddleware())
	e.PUT("/projects/:id", ph.PutProjectHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/projects/:id", ph.DeleteProjectHandler(), _middlewares.JWTMiddleware())

	e.POST("/task", th.PostTaskHandler(), _middlewares.JWTMiddleware())
	e.GET("/task", th.GetAllHandler(), _middlewares.JWTMiddleware())
	e.PUT("/task/:id", th.PutTaskHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/task/:id", th.DeleteTaskHandler(), _middlewares.JWTMiddleware())
	e.POST("/task/:id/completed", th.PostTaskCompleteHandler(), _middlewares.JWTMiddleware())
	e.POST("/task/:id/reopen", th.PostTaskReOpenHandler(), _middlewares.JWTMiddleware())

}
