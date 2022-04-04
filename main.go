package main

import (
	"fmt"
	"log"
	"project2/configs"
	_userHandler "project2/delivery/handler/user"
	_routes "project2/delivery/routes"
	_userRepository "project2/repository/user"
	_userUseCase "project2/usecase/user"
	_utils "project2/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()

	_routes.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
