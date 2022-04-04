package routes

import (
	_userHandler "project2/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.PostUserHandler())
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetUserHandler())
	e.PUT("/users/:id", uh.PutUserHandler())
	e.DELETE("/users/:id", uh.DeleteUserHandler())
}
