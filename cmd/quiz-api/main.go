package main

import (
	"quiz-api/config"
	"quiz-api/controller"
	"quiz-api/database"
	"quiz-api/repository"
	"quiz-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config := config.New()
	db := database.ConnectDatabase(config)
	repos := repository.InitRepositories(db, config)
	controllers := controller.InitControllers(repos)

	routes.Routes(e, controllers)

	e.Logger.Fatal(e.Start(config.Listen_Addr))

}
