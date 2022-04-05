package routes

import (
	_authHandler "project2/delivery/handler/auth"
	_projectHandler "project2/delivery/handler/project"
	_middlewares "project2/delivery/middlewares"

	_userHandler "project2/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, ph *_projectHandler.ProjectHandler) {
	e.POST("/users", uh.PostUserHandler())
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.PutUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())

	e.POST("/projects", ph.PostProjectHandler(), _middlewares.JWTMiddleware())
	e.GET("/Projects", ph.GetAllHandler(), _middlewares.JWTMiddleware())
	e.PUT("/projects/:id", ph.PutProjectHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/Projects/:id", ph.DeleteProjectHandler(), _middlewares.JWTMiddleware())
}
