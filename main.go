package main

import (
	"fmt"
	"log"
	"project2/configs"
	_authHandler "project2/delivery/handler/auth"
	_userHandler "project2/delivery/handler/user"
	_middlewares "project2/delivery/middlewares"
	_routes "project2/delivery/routes"
	_authRepository "project2/repository/auth"
	_userRepository "project2/repository/user"
	_authUseCase "project2/usecase/auth"
	_userUseCase "project2/usecase/user"
	_utils "project2/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
